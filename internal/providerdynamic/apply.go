package providerdynamic

import (
	"context"
	"fmt"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
)

type sobject struct {
	sobjects.BaseSObject
	typ string
}

func (o sobject) ApiName() string {
	return o.typ
}

// ApplyResourceChange function
func (s *providerServer) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	resp := &tfprotov5.ApplyResourceChangeResponse{}
	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine planned resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	applyPlannedState, err := req.PlannedState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal planned resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ApplyResourceChange]", "[PlannedState]", spew.Sdump(applyPlannedState))

	applyPlannedVal := make(map[string]tftypes.Value)
	err = applyPlannedState.As(&applyPlannedVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract planned resource state from tftypes.Value",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	applyPriorState, err := req.PriorState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal prior resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ApplyResourceChange]", "[PriorState]", spew.Sdump(applyPriorState))

	applyPriorVal := make(map[string]tftypes.Value)
	err = applyPriorState.As(&applyPriorVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract planned resource state from tftypes.Value",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	switch {
	case !applyPlannedState.IsNull():
		// Apply resource

		var typ string
		err = applyPlannedVal["type"].As(&typ)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to extract 'type' from planned state",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		attrs, err := mapOfInterfaceFromTftypesObject(applyPlannedVal["attributes"])
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Error extracting 'attributes' from planned state",
				Detail:   err.Error(),
			})
			return resp, nil
		}
		forceResp := &force.SObjectResponse{}
		meta, err := s.client.DescribeSObjects()
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Error getting sObject metadata",
				Detail:   err.Error(),
			})
			return resp, nil
		}
		uri := meta[typ].URLs["sobject"]

		err = s.client.Post(uri, nil, attrs, &forceResp)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Error creating sObject: " + typ,
				Detail:   err.Error(),
			})
			return resp, nil
		}

		// we should probably do a full read, but let's just see if setting the ID back
		// in state works
		applyPlannedVal["id"] = tftypes.NewValue(tftypes.String, forceResp.Id)

		newStateVal := tftypes.NewValue(applyPlannedState.Type(), applyPlannedVal)
		s.logger.Trace("[ApplyResourceChange][Apply]", "new state value", spew.Sdump(newStateVal))

		newResState, err := tfprotov5.NewDynamicValue(newStateVal.Type(), newStateVal)
		if err != nil {
			return resp, err
		}

		resp.NewState = &newResState

	case applyPlannedState.IsNull():
		// Delete the resource
		// PoC focuses on User, which is a special case that can't be deleted
		// we can just deactivate it and throw away state
		// There will be a clear need to override select operations based on the sObject
		// type.

		var id string
		err = applyPriorVal["id"].As(&id)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to extract 'id' from prior state",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		var typ string
		err = applyPriorVal["type"].As(&typ)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to extract 'type' from prior state",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		type dynamic struct {
			sobject
			IsActive bool
		}

		err = s.client.UpdateSObject(id, &dynamic{
			sobject: sobject{
				typ: typ,
			},
			IsActive: false,
		})
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Error updating sObject: " + typ,
				Detail:   err.Error(),
			})
			return resp, nil
		}

		resp.NewState = req.PlannedState
	}

	return resp, nil
}

func mapOfInterfaceFromTftypesObject(in tftypes.Value) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	attrs := make(map[string]tftypes.Value)

	if err := in.As(&attrs); err != nil {
		return result, err
	}

	for key, value := range attrs {
		if !value.IsNull() {
			switch {
			case value.Type().Is(tftypes.Bool):
				var b bool
				if err := value.As(&b); err != nil {
					return result, err
				}
				result[key] = b
			case value.Type().Is(tftypes.String):
				var s string
				if err := value.As(&s); err != nil {
					return result, err
				}
				result[key] = s
			case value.Type().Is(tftypes.Number):
				// because value.IsNull() == false we can be sure
				// if the value is still 0, after the .As() call, it's
				// because it's actually set to 0, so the starting value
				// in NewFloat() really doesn't matter
				f := big.NewFloat(0)
				if err := value.As(f); err != nil {
					return result, err
				}
				ff, _ := f.Float64()
				// TODO do we care about accuracy?
				result[key] = ff
			default:
				// PoC focuses on sObject which is flat structure,
				// should be representable by primitive JSON type
				return result, fmt.Errorf("Unexpected type: %v", value)
			}
		}
	}
	return result, nil
}
