package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
)

type profileType struct {
}

func (profileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"user_license_id": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
		},
	}, nil
}

func (p profileType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(p)}
	}
	return profileResource{client: provider.client}, nil
}

type profile struct {
	Id            types.String `tfsdk:"id" force:"-"`
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   *string      `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
}

func (profile) ApiName() string {
	return "Profile"
}

func (profile) ExternalIdApiName() string {
	return ""
}

// Zero out fields that can't be updated.
// Salesforce doesn't compare with the existing data, it rejects the request if
// the field is present. Ensure the fields have the omitempty tag for this to work.
func (p profile) updatable() profile {
	p.UserLicenseId = ""
	return p
}

type profileResource struct {
	client *force.ForceApi
}

func (p profileResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var prof profile
	if diags := req.Plan.Get(ctx, &prof); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := p.client.InsertSObject(prof)
	if err != nil {
		resp.AddError("Error inserting Profile", err.Error())
		return
	}
	prof.Id = types.String{Value: sfResp.Id}

	resp.Diagnostics = resp.State.Set(ctx, &prof)
}

func (p profileResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var prof profile
	if diags := req.State.Get(ctx, &prof); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := p.client.GetSObject(prof.Id.Value, nil, &prof)
	if err != nil {
		if isErrorNotFound(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.AddError("Error getting Profile", err.Error())
		}
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &prof)
}

func (p profileResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var prof profile
	if diags := req.Plan.Get(ctx, &prof); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := p.client.UpdateSObject(prof.Id.Value, prof.updatable())
	if err != nil {
		resp.AddError("Error updating Profile", err.Error())
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &prof)
}

func (p profileResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var prof profile
	if diags := req.State.Get(ctx, &prof); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := p.client.DeleteSObject(prof.Id.Value, prof)
	if err != nil {
		if !isErrorNotFound(err) {
			resp.AddError("Error deleting Profile", err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}
