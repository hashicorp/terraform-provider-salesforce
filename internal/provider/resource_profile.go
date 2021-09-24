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
					NormalizeId{},
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
	return &profileResource{
		Resource: Resource{
			Client: provider.client,
			Data:   &profileResourceData{},
		},
	}, nil
}

type profileResource struct {
	Resource
}

type profileResourceData struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   *string      `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
	Id            types.String `tfsdk:"id" force:"-"`
}

func (profileResourceData) ApiName() string {
	return "Profile"
}

func (profileResourceData) ExternalIdApiName() string {
	return ""
}

func (p *profileResourceData) Instance() force.SObject {
	return p
}

func (p *profileResourceData) Updatable() force.SObject {
	updatable := *p
	updatable.UserLicenseId = ""
	return updatable
}

func (p *profileResourceData) GetId() string {
	return p.Id.Value
}

func (p *profileResourceData) SetId(id string) {
	p.Id = types.String{Value: id}
}
