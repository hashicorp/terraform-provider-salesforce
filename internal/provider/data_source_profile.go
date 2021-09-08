package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
)

type profileType struct {
}

func (p profileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
		},
	}, nil
}

func (p profileType) NewDataSource(_ context.Context, prov tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{
			diag.NewErrorDiagnostic("Error converting provider", fmt.Sprintf("An unexpected error was encountered converting the provider. This is always a bug in the provider.\n\nType: %T", p)),
		}
	}
	return profileDataSource{client: provider.client}, nil
}

type profileDataSource struct {
	client *force.ForceApi
}

type profileData struct {
	ID   types.String `tfsdk:"id"`
	Name string       `tfsdk:"name"`
}

type Profile struct {
	sobjects.BaseSObject
}

type ProfileQueryResponse struct {
	sobjects.BaseQuery
	Records []Profile `json:"Records" force:"records"`
}

func (p profileDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var pData profileData
	if diags := req.Config.Get(ctx, &pData); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	var query ProfileQueryResponse
	nameFilter := fmt.Sprintf("Name = '%s'", pData.Name)
	if err := p.client.Query(force.BuildQuery("Id, Name", "Profile", []string{nameFilter}), &query); err != nil {
		resp.Diagnostics.AddError("Error getting profile", err.Error())
		return
	}
	profile := query.Records[0]
	pData.ID = types.String{Value: profile.Id}

	if diags := resp.State.Set(ctx, &pData); diags.HasError() {
		resp.Diagnostics = diags
	}
}
