package provider

import (
	"context"
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
	Name          string       `tfsdk:"name"`
	Description   *string      `tfsdk:"description"`
	UserLicenseId string       `tfsdk:"user_license_id"`
	Permissions   types.Map    `tfsdk:"permissions"`
	Id            types.String `tfsdk:"id"`
}

func (p profileResourceData) PermissionKeys(prefix string) []string {
	var keys []string
	for k := range p.Permissions.Elems {
		keys = append(keys, prefix+k)
	}
	return keys
}

func (p profileResourceData) ToMap(exclude ...string) profileMap {
	pMap := make(profileMap)
	if p.Name != "" {
		pMap["Name"] = p.Name
	}
	if p.Description != nil && *p.Description != "" {
		pMap["Description"] = *p.Description
	}
	if p.UserLicenseId != "" {
		pMap["UserLicenseId"] = p.UserLicenseId
	}
	// flatten permissions
	for k, v := range p.Permissions.Elems {
		key := "Permissions" + k
		pMap[key] = v.(types.Bool).Value
	}
	// exclude keys, useful for update
	for _, k := range exclude {
		delete(pMap, k)
	}
	return pMap
}

// due to the permissions map, we need a separate type that is flattened to
// what SF expects
type profileMap map[string]interface{}

func (p profileMap) ToStateData(includePermissions ...string) profileResourceData {
	data := profileResourceData{
		Name:          p["Name"].(string),
		UserLicenseId: p["UserLicenseId"].(string),
	}
	desc := p["Description"].(string)
	data.Description = &desc
	// expand permissions
	permissions := make(map[string]attr.Value, len(includePermissions))
	for _, k := range includePermissions {
		v, ok := p[k]
		trimmedKey := strings.TrimPrefix(k, "Permissions")
		if ok {
			permissions[trimmedKey] = types.Bool{Value: v.(bool)}
		} else {
			// set to unknown, maybe we should panic?
			permissions[trimmedKey] = types.Bool{Unknown: true}
		}
	}
	if len(permissions) > 0 {
		data.Permissions = types.Map{ElemType: types.BoolType, Elems: permissions}
	}
	return data
}

func (profileMap) ApiName() string {
	return "Profile"
}

func (profileMap) ExternalIdApiName() string {
	return ""
}

func (p *profileResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data profileResourceData
	if diags := req.Plan.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := p.client.InsertSObject(data.ToMap())
	if err != nil {
		resp.Diagnostics.AddError("Error Inserting Profile", err.Error())
		return
	}
	data.Id = types.String{Value: sfResp.Id}

	resp.Diagnostics = resp.State.Set(ctx, &data)
}

func (p *profileResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data profileResourceData
	if diags := req.State.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	var pMap profileMap
	if err := p.client.GetSObject(data.Id.Value, nil, &pMap); err != nil {
		if isNotFoundError(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error Getting Profile", err.Error())
		}
		return
	}

	d := pMap.ToStateData(data.PermissionKeys("Permissions")...)
	// copy the ID back over
	d.Id = data.Id

	resp.Diagnostics = resp.State.Set(ctx, &d)
}

func (p *profileResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data profileResourceData
	if diags := req.Plan.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	if err := p.client.UpdateSObject(data.Id.Value, data.ToMap("UserLicenseId")); err != nil {
		resp.Diagnostics.AddError("Error Updating Profile", err.Error())
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &data)
}

func (p *profileResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data profileResourceData
	if diags := req.State.Get(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	if err := p.client.DeleteSObject(data.Id.Value, data.ToMap()); err != nil {
		if !isNotFoundError(err) {
			resp.Diagnostics.AddError("Error Deleting Profile", err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

func (p *profileResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	id := normalizeId(req.ID)
	var pMap profileMap
	if err := p.client.GetSObject(id, nil, &pMap); err != nil {
		resp.Diagnostics.AddError("Error Importing Profile", err.Error())
		return
	}
	data := pMap.ToStateData()
	data.Id = types.String{Value: id}
	if diags := resp.State.Set(ctx, &data); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	resp.Diagnostics.AddWarning("Profile imported without permissions", "permissions can be explicitly set with the permissions = {} attribute after import, but existing permission settings cannot be imported due to technical limitations.")
}
