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
			"developer_name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				}, // TODO full validation, see requirements in https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_role.htm
			},
			"parent_role_id": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (u userRoleType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(u)}
	}
	return userRoleResource{client: provider.client}, nil
}

type userRole struct {
	Id            types.String `tfsdk:"id" force:"-"`
	Name          string       `tfsdk:"name" force:",omitempty"`
	DeveloperName string       `tfsdk:"developer_name" force:",omitempty"`
	ParentRoleId  *string      `tfsdk:"parent_role_id" force:",omitempty"`
}

func (userRole) ApiName() string {
	return "UserRole"
}

func (userRole) ExternalIdApiName() string {
	return ""
}

type userRoleResource struct {
	client *force.ForceApi
}

func (u userRoleResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var role userRole
	if diags := req.Plan.Get(ctx, &role); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := u.client.InsertSObject(role)
	if err != nil {
		resp.AddError("Error inserting UserRole", err.Error())
		return
	}
	role.Id = types.String{Value: sfResp.Id}

	resp.Diagnostics = resp.State.Set(ctx, &role)
}

func (u userRoleResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var role userRole
	if diags := req.State.Get(ctx, &role); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := u.client.GetSObject(role.Id.Value, nil, &role)
	if err != nil {
		if isErrorNotFound(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.AddError("Error getting UserRole", err.Error())
		}
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &role)
}

func (u userRoleResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var role userRole
	if diags := req.Plan.Get(ctx, &role); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := u.client.UpdateSObject(role.Id.Value, role)
	if err != nil {
		resp.AddError("Error updating UserRole", err.Error())
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &role)
}

func (u userRoleResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var role userRole
	if diags := req.State.Get(ctx, &role); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := u.client.DeleteSObject(role.Id.Value, role)
	if err != nil {
		if !isErrorNotFound(err) {
			resp.AddError("Error deleting UserRole", err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}
