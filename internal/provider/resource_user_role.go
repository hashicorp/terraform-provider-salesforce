package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
)

type userRoleType struct {
}

func (userRoleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "ID of the resource.",
				Type:        types.StringType,
				Computed:    true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					staticComputed{},
				},
			},
			"name": {
				Description: "Name of the role. Corresponds to Label on the user interface.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
				},
			},
			"developer_name": {
				Description: "The unique name of the object in the API. This name can contain only underscores and alphanumeric characters, and must be unique in your org. It must begin with a letter, not include spaces, not end with an underscore, and not contain two consecutive underscores. In managed packages, this field prevents naming conflicts on package installations. With this field, a developer can change the object’s name in a managed package and the changes are reflected in a subscriber’s organization. Corresponds to Role Name in the user interface.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					// TODO full validation, see requirements in https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_role.htm
					/*
						Developer Name: The User Role API Name can only contain underscores and alphanumeric characters.
						It must be unique, begin with a letter, not include spaces, not end with an
						underscore, and not contain two consecutive underscores.
					*/
					notEmptyString{},
				},
			},
			"parent_role_id": {
				Description: "The ID of the parent role.",
				Type:        types.StringType,
				Optional:    true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
					fixNullToUnknown{},
				},
			},
		},
	}, nil
}

func (u userRoleType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(u)}
	}
	return &userRoleResource{
		Resource: Resource{
			Client: provider.client,
			Data:   &userRoleResourceData{},
		},
	}, nil
}

type userRoleResource struct {
	Resource
}

type userRoleResourceData struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	DeveloperName string       `tfsdk:"developer_name" force:",omitempty"`
	ParentRoleId  *string      `tfsdk:"parent_role_id" force:",omitempty"`
	Id            types.String `tfsdk:"id" force:"-"`
}

func (userRoleResourceData) ApiName() string {
	return "UserRole"
}

func (userRoleResourceData) ExternalIdApiName() string {
	return ""
}

func (u *userRoleResourceData) Instance() force.SObject {
	return u
}

func (u *userRoleResourceData) Insertable() force.SObject {
	return *u
}

func (u *userRoleResourceData) Updatable() force.SObject {
	return *u
}

func (u *userRoleResourceData) GetId() string {
	return u.Id.Value
}

func (u *userRoleResourceData) SetId(id string) {
	u.Id = types.String{Value: id}
}
