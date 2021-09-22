package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/nimajalali/go-force/force"
)

type SObject interface {
	Instance() force.SObject
	Updatable() force.SObject
	SetId(string)
	GetId() string
}

type Resource struct {
	Client  *force.ForceApi `force:"-"`
	SObject `force:"-"`
}

func (r Resource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	sobject := r.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := r.Client.InsertSObject(sobject)
	if err != nil {
		resp.AddError(fmt.Sprintf("Error inserting %s", sobject.ApiName()), err.Error())
		return
	}
	r.SetId(sfResp.Id)

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r Resource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	sobject := r.Instance()
	if diags := req.State.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.GetSObject(r.GetId(), nil, sobject)
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

func (r Resource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	sobject := r.Instance()
	if diags := req.Plan.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.UpdateSObject(r.GetId(), r.Updatable())
	if err != nil {
		resp.AddError(fmt.Sprintf("Error updating %s", sobject.ApiName()), err.Error())
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, sobject)
}

func (r Resource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	sobject := r.Instance()
	if diags := req.State.Get(ctx, sobject); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	err := r.Client.DeleteSObject(r.GetId(), sobject)
	if err != nil {
		if !isErrorNotFound(err) {
			resp.AddError(fmt.Sprintf("Error deleting %s", sobject.ApiName()), err.Error())
			return
		}
	}

	resp.State.RemoveResource(ctx)
}
