package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/nimajalali/go-force/force"
)

type ResourceData interface {
	Instance() force.SObject
	Insertable() force.SObject
	Updatable() force.SObject
	SetId(string)
	GetId() string
}

type Resource struct {
	Client              *force.ForceApi
	Data                ResourceData
	NeedsGetAfterUpsert bool
}

func (r *Resource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	sobject := r.Data.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := r.Client.InsertSObject(r.Data.Insertable())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Inserting %s", sobject.ApiName()), err.Error())
		return
	}
	r.Data.SetId(sfResp.Id)

	if r.NeedsGetAfterUpsert {
		err := r.Client.GetSObject(r.Data.GetId(), nil, sobject)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", sobject.ApiName()), err.Error())
			return
		}
	}

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
			resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", sobject.ApiName()), err.Error())
		}
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r *Resource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	sobject := r.Data.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.UpdateSObject(r.Data.GetId(), r.Data.Updatable())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Updating %s", sobject.ApiName()), err.Error())
		return
	}

	if r.NeedsGetAfterUpsert {
		err := r.Client.GetSObject(r.Data.GetId(), nil, sobject)
		if err != nil {
			if isErrorNotFound(err) {
				resp.State.RemoveResource(ctx)
			} else {
				resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", sobject.ApiName()), err.Error())
			}
			return
		}
	}

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
			resp.Diagnostics.AddError(fmt.Sprintf("Error Deleting %s", sobject.ApiName()), err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

func (r *Resource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	sobject := r.Data.Instance()
	id := normalizeId(req.ID)
	err := r.Client.GetSObject(id, nil, sobject)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Getting %s", sobject.ApiName()), err.Error())
		return
	}
	r.Data.SetId(id)

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}
