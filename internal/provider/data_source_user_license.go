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

type userLicenseDatasourceType struct {
}

func (userLicenseDatasourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"license_definition_key": {
				Type:     types.StringType,
				Required: true,
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

type UserLicense struct {
	ID                   types.String `tfsdk:"id" force:"-"` // go-force can't serialize nonprimitives
	Id                   string       `tfsdk:"-" force:",omitempty"`
	LicenseDefinitionKey string       `tfsdk:"license_definition_key"`
}

func (UserLicense) ApiName() string {
	return "UserLicense"
}

func (UserLicense) ExternalIdApiName() string {
	return ""
}

type UserLicenseQueryResponse struct {
	sobjects.BaseQuery
	Records []UserLicense `json:"Records" force:"records"`
}

func (u userLicenceDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var userLicense UserLicense
	if diags := req.Config.Get(ctx, &userLicense); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	var query UserLicenseQueryResponse
	licenseDefinitionKeyFilter := fmt.Sprintf("LicenseDefinitionKey = '%s'", userLicense.LicenseDefinitionKey)
	if err := u.client.Query(force.BuildQuery("Id, LicenseDefinitionKey", userLicense.ApiName(), []string{licenseDefinitionKeyFilter}), &query); err != nil {
		resp.Diagnostics.AddError("Error getting User License", err.Error())
		return
	}
	userLicense = query.Records[0]
	userLicense.ID = types.String{Value: userLicense.Id}

	resp.Diagnostics = resp.State.Set(ctx, &userLicense)
}
