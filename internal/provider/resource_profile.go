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
			},
			"description": {
				Type:     types.StringType,
				Required: true,
			},
			"user_license_id": {
				Type:     types.StringType,
				Required: true,
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

type Profile struct {
	ID            types.String `tfsdk:"id" force:"-"`
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   string       `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
}

func (Profile) ApiName() string {
	return "Profile"
}

func (Profile) ExternalIdApiName() string {
	return ""
}

type profileResource struct {
	client *force.ForceApi
}

func (p profileResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var profile Profile
	if diags := req.Plan.Get(ctx, &profile); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := p.client.InsertSObject(profile)
	if err != nil {
		resp.AddError("Error inserting Profile", err.Error())
		return
	}
	profile.ID = types.String{Value: sfResp.Id}

	resp.Diagnostics = resp.State.Set(ctx, &profile)
}

func (p profileResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
}

func (p profileResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

func (p profileResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}
