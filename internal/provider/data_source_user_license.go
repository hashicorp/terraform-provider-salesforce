package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-salesforce/internal/picklists"
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
)

type userLicenseDatasourceType struct {
}

func (userLicenseDatasourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "ID of the resource.",
				Type:        types.StringType,
				Computed:    true,
			},
			"license_definition_key": {
				Description: "A string that uniquely identifies a particular user license. Valid options vary depending on organization type and configuration. For a complete list see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_userlicense.htm",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					stringInSlice{slice: picklists.LicenseDefinitionKeys},
				},
			},
		},
	}, nil
}

func (u userLicenseDatasourceType) NewDataSource(_ context.Context, prov tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(u)}
	}
	return userLicenceDataSource{client: provider.client}, nil
}

type userLicenceDataSource struct {
	client *force.ForceApi
}

type userLicenseData struct {
	Id                   *string `tfsdk:"id"`
	LicenseDefinitionKey string  `tfsdk:"license_definition_key"`
}

type userLicenseQueryResponse struct {
	sobjects.BaseQuery
	Records []userLicenseData
}

func (u userLicenceDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var uData userLicenseData
	if diags := req.Config.Get(ctx, &uData); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	var query userLicenseQueryResponse
	licenseDefinitionKeyFilter := fmt.Sprintf("LicenseDefinitionKey = '%s'", uData.LicenseDefinitionKey)
	if err := u.client.Query(force.BuildQuery("Id, LicenseDefinitionKey", "UserLicense", []string{licenseDefinitionKeyFilter}), &query); err != nil {
		resp.Diagnostics.AddError("Error Getting User License", err.Error())
		return
	}
	if len(query.Records) == 0 {
		resp.Diagnostics.AddError("Error Getting User License", fmt.Sprintf("No User License where %s", licenseDefinitionKeyFilter))
		return
	}

	uData = query.Records[0]
	resp.Diagnostics = resp.State.Set(ctx, &uData)
}
