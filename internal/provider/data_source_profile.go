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

type profileDatasourceType struct {
}

func (profileDatasourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

func (p profileDatasourceType) NewDataSource(_ context.Context, prov tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(p)}
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

type ProfileQueryResponse struct {
	sobjects.BaseQuery
	Records []struct {
		sobjects.BaseSObject
	} `json:"Records" force:"records"`
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
