package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
				PlanModifiers: tfsdk.AttributePlanModifiers{
					staticComputed{},
				},
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
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
					notEmptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
					tfsdk.RequiresReplace(),
				},
			},
			"permissions": {
				Type:     types.MapType{ElemType: types.BoolType},
				Optional: true,
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
		client: provider.client,
	}, nil
}

type profileResource struct {
	client *force.ForceApi
}

type profileResourceData struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   *string      `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
	Permissions   types.Map    `tfsdk:"permissions" force:"-"`
	Id            types.String `tfsdk:"id" force:"-"`
}

func (profileResourceData) ApiName() string {
	return "Profile"
}

func (profileResourceData) ExternalIdApiName() string {
	return ""
}

func (p *profileResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data profileResourceData
	if diags := req.Plan.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	// create a resource without any permissions
	sfResp, err := p.client.InsertSObject(data)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Inserting %s", data.ApiName()), err.Error())
		return
	}
	data.Id = types.String{Value: sfResp.Id}

	perms := permissionsMapFromTypesMap(data.Permissions.Elems)
	if len(perms) > 0 {
		if err := p.client.UpdateSObject(data.Id.Value, perms); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Updating %s Permissions", data.ApiName()), err.Error())
			return
		}
	}

	if err := p.client.GetSObject(data.Id.Value, nil, &data); err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", data.ApiName()), err.Error())
		return
	}

	if len(perms) > 0 {
		// we re-fetch the entire resource into a generic map[string]interface{}
		var profile permissionsMap
		if err := p.client.GetSObject(data.Id.Value, nil, &profile); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s Permissions", data.ApiName()), err.Error())
			return
		}
		// and just prune down to the user specified ones
		data.Permissions = types.Map{ElemType: types.BoolType, Elems: mergePermissions(perms, profile)}
	}

	resp.Diagnostics = resp.State.Set(ctx, &data)
}

func (p *profileResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data profileResourceData
	if diags := req.State.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	if err := p.client.GetSObject(data.Id.Value, nil, &data); err != nil {
		if isNotFoundError(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", data.ApiName()), err.Error())
		}
		return
	}

	if len(data.Permissions.Elems) > 0 {
		var profile permissionsMap
		if err := p.client.GetSObject(data.Id.Value, nil, &profile); err != nil {
			if isNotFoundError(err) {
				resp.State.RemoveResource(ctx)
			} else {
				resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s Permissions", data.ApiName()), err.Error())
			}
			return
		}
		data.Permissions = types.Map{ElemType: types.BoolType, Elems: mergePermissions(permissionsMapFromTypesMap(data.Permissions.Elems), profile)}
	}

	resp.Diagnostics = resp.State.Set(ctx, &data)
}

func (p *profileResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data profileResourceData
	if diags := req.Plan.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	// zero out fields that can't be updated
	updatable := data
	updatable.UserLicenseId = ""

	if err := p.client.UpdateSObject(data.Id.Value, updatable); err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Updating %s", data.ApiName()), err.Error())
		return
	}

	perms := permissionsMapFromTypesMap(data.Permissions.Elems)
	if len(perms) > 0 {
		if err := p.client.UpdateSObject(data.Id.Value, perms); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Updating %s Permissions", data.ApiName()), err.Error())
			return
		}
	}

	if err := p.client.GetSObject(data.Id.Value, nil, &data); err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", data.ApiName()), err.Error())
		return
	}

	if len(perms) > 0 {
		// we re-fetch the entire resource into a generic map[string]interface{}
		var profile permissionsMap
		if err := p.client.GetSObject(data.Id.Value, nil, &profile); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s Permissions", data.ApiName()), err.Error())
			return
		}
		// and just prune down to the user specified ones
		data.Permissions = types.Map{ElemType: types.BoolType, Elems: mergePermissions(perms, profile)}
	}

	resp.Diagnostics = resp.State.Set(ctx, &data)
}

func (p *profileResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data profileResourceData
	if diags := req.State.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	if err := p.client.DeleteSObject(data.Id.Value, data); err != nil {
		if !isNotFoundError(err) {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Deleting %s", data.ApiName()), err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

func (p *profileResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	var data profileResourceData
	id := normalizeId(req.ID)
	if err := p.client.GetSObject(id, nil, &data); err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", data.ApiName()), err.Error())
		return
	}
	data.Id = types.String{Value: id}
	if diags := resp.State.Set(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	resp.Diagnostics.AddWarning("Profile imported without permissions", "specific permissions can always be explicitly set with the permissions = {} attribute after import, but existing permission settings cannot be imported due to technical limitations.")
}

func permissionsMapFromTypesMap(m map[string]attr.Value) permissionsMap {
	permissions := make(permissionsMap, len(m))
	for k, v := range m {
		val := v.(types.Bool)
		if !val.Unknown && !val.Null {
			// SF API has a Permissions prefix to all the permissions attributes
			key := "Permissions" + k
			permissions[key] = val.Value
		}
	}
	return permissions
}

func mergePermissions(sent, response map[string]interface{}) map[string]attr.Value {
	m := make(map[string]attr.Value, len(sent))
	for k := range sent {
		// remove the redundant prefix when saving into state
		key := strings.TrimPrefix(k, "Permissions")
		if v, ok := response[k]; ok {
			m[key] = types.Bool{Value: v.(bool)}
		} else {
			// set to unknown to force an error, this should not happen
			m[key] = types.Bool{Unknown: true}
		}
	}
	return m
}

type permissionsMap map[string]interface{}

func (permissionsMap) ApiName() string {
	return "Profile"
}

func (permissionsMap) ExternalIdApiName() string {
	return ""
}
