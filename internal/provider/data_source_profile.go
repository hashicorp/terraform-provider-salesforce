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
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
				},
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
	Id   *string `tfsdk:"id"`
	Name string  `tfsdk:"name"`
}

type profileQueryResponse struct {
	sobjects.BaseQuery
	Records []profileData
}

func (p profileDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var pData profileData
	if diags := req.Config.Get(ctx, &pData); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	var query profileQueryResponse
	nameFilter := fmt.Sprintf("Name = '%s'", pData.Name)
	if err := p.client.Query(force.BuildQuery("Id, Name", "Profile", []string{nameFilter}), &query); err != nil {
		resp.Diagnostics.AddError("Error getting Profile", err.Error())
		return
	}
	if len(query.Records) == 0 {
		resp.Diagnostics.AddError("Error getting Profile", fmt.Sprintf("No Profile where %s", nameFilter))
		return
	}

	pData = query.Records[0]
	resp.Diagnostics = resp.State.Set(ctx, &pData)
}
