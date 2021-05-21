package providerdynamic

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/nimajalali/go-force/force"
)

// ReadResource function
func (s *providerServer) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	resp := &tfprotov5.ReadResourceResponse{}
	var resState map[string]tftypes.Value
	var err error
	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	currentState, err := req.CurrentState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to decode current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	if currentState.IsNull() {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to read resource",
			Detail:   "Incomplete of missing state",
		})
		return resp, nil
	}
	err = currentState.As(&resState)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract resource from current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var id string
	err = resState["id"].As(&id)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract 'id' from current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var typ string
	err = resState["type"].As(&typ)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract 'type' from current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	sObject, err := getSObject(id, typ, s.client)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get sObject",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	resState["attributes"] = tftypes.NewValue(tftypes.Object{
		AttributeTypes: attributeTypes(sObject),
	}, sObject)

	nsVal := tftypes.NewValue(currentState.Type(), resState)
	newState, err := tfprotov5.NewDynamicValue(nsVal.Type(), nsVal)
	if err != nil {
		return resp, err
	}
	resp.NewState = &newState
	return resp, nil
}

func attributeTypes(values map[string]tftypes.Value) map[string]tftypes.Type {
	types := make(map[string]tftypes.Type, len(values))
	for attr, value := range values {
		types[attr] = value.Type()
	}
	return types
}

func getSObject(id string, typ string, client *force.ForceApi) (map[string]tftypes.Value, error) {
	result := make(map[string]tftypes.Value)
	meta, err := client.DescribeSObjects()
	if err != nil {
		return result, err
	}
	uri := strings.Replace(meta[typ].URLs["rowTemplate"], "{ID}", id, 1)
	out := make(map[string]interface{})
	err = client.Get(uri, nil, &out)
	if err != nil {
		return result, err
	}

	for key, value := range out {
		if key == "attributes" {
			// this is a special part of the sObject response that isn't useful
			continue
		}
		if key == "Id" {
			// this is redundant
			continue
		}
		switch v := value.(type) {
		case nil:
			// do nothing
		case float64:
			result[key] = tftypes.NewValue(tftypes.Number, v)
		case string:
			result[key] = tftypes.NewValue(tftypes.String, v)
		case bool:
			result[key] = tftypes.NewValue(tftypes.Bool, v)
		default:
			return result, fmt.Errorf("Invalid type key: %s, val: %v", key, v)
		}
	}

	return result, nil
}
