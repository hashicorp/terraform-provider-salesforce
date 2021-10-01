package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/nimajalali/go-force/force"
)

type ResourceData interface {
	Instance() force.SObject
	Updatable() force.SObject
	SetId(string)
	GetId() string
}

type Resource struct {
	Client *force.ForceApi
	Data   ResourceData
}

func (r *Resource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	sobject := r.Data.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := r.Client.InsertSObject(sobject)
	if err != nil {
		resp.AddError(fmt.Sprintf("Error inserting %s", sobject.ApiName()), err.Error())
		return
	}
	r.Data.SetId(sfResp.Id)

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r *Resource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	sobject := r.Data.Instance()
	if diags := req.State.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.GetSObject(r.Data.GetId(), nil, sobject)
	if err != nil {
		if isErrorNotFound(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.AddError(fmt.Sprintf("Error getting %s", sobject.ApiName()), err.Error())
		}
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r *Resource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// id, diags := req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"))
	// if diags.HasError() {
	// 	resp.Diagnostics = diags
	// 	return
	// }
	// idStr := id.(types.String).Value
	sobject := r.Data.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.UpdateSObject(r.Data.GetId(), r.Data.Updatable())
	if err != nil {
		resp.AddError(fmt.Sprintf("Error updating %s", sobject.ApiName()), err.Error())
		return
	}
	//r.Data.SetId(idStr)

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r *Resource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	sobject := r.Data.Instance()
	if diags := req.State.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.DeleteSObject(r.Data.GetId(), sobject)
	if err != nil {
		if !isErrorNotFound(err) {
			resp.AddError(fmt.Sprintf("Error deleting %s", sobject.ApiName()), err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

func (r *Resource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
