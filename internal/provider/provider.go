// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-salesforce/internal/auth"
	"github.com/nimajalali/go-force/force"
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	client *force.ForceApi
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "A Provider for managing a Salesforce Organization",
		Attributes: map[string]tfsdk.Attribute{
			"client_id": {
				Description: "Client ID of the connected app. Corresponds to Consumer Key in the user interface. Can be specified with the environment variable SALESFORCE_CLIENT_ID.",
				Type:        types.StringType,
				Optional:    true,
			},
			"private_key": {
				Description: "Private Key associated to the public certificate that was uploaded to the connected app. This may point to a file location or be set directly. This should not be confused with the Consumer Secret in the user interface. Can be specified with the environment variable SALESFORCE_PRIVATE_KEY.",
				Type:        types.StringType,
				Optional:    true,
				Sensitive:   true,
			},
			"api_version": {
				Description: "API version of the salesforce org in the format in the format: MAJOR.MINOR (please omit any leading 'v'). The provider requires at least version 53.0. Can be specified with the environment variable SALESFORCE_API_VERSION.",
				Type:        types.StringType,
				Optional:    true,
			},
			"username": {
				Description: "Salesforce Username of a System Administrator like user for the provider to authenticate as. Can be specified with the environment variable SALESFORCE_USERNAME.",
				Type:        types.StringType,
				Optional:    true,
			},
			"login_url": {
				Description: "Directs the authentication request, defaults to the production endpoint https://login.salesforce.com, should be set to https://test.salesforce.com for sandbox organizations. Can be specified with the environment variable SALESFORCE_LOGIN_URL.",
				Type:        types.StringType,
				Optional:    true,
			},
		},
	}, nil
}

type providerData struct {
	ClientId   types.String `tfsdk:"client_id"`
	PrivateKey types.String `tfsdk:"private_key"`
	ApiVersion types.String `tfsdk:"api_version"`
	Username   types.String `tfsdk:"username"`
	LoginUrl   types.String `tfsdk:"login_url"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var config providerData
	if diags := req.Config.Get(ctx, &config); diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	// interpolation not allowed in provider block
	if config.ClientId.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "client_id")
		return
	}
	if config.PrivateKey.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "private_key")
		return
	}
	if config.ApiVersion.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "api_version")
		return
	}
	if config.Username.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "username")
		return
	}
	if config.LoginUrl.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "login_url")
		return
	}

	// if unset, fallback to env
	if config.ClientId.Null {
		config.ClientId.Value = os.Getenv("SALESFORCE_CLIENT_ID")
	}
	if config.PrivateKey.Null {
		config.PrivateKey.Value = os.Getenv("SALESFORCE_PRIVATE_KEY")
	}
	if config.ApiVersion.Null {
		config.ApiVersion.Value = os.Getenv("SALESFORCE_API_VERSION")
	}
	if config.Username.Null {
		config.Username.Value = os.Getenv("SALESFORCE_USERNAME")
	}
	if config.LoginUrl.Null {
		config.LoginUrl.Value = os.Getenv("SALESFORCE_LOGIN_URL")
	}

	// required if still unset
	if config.ClientId.Value == "" {
		addAttributeMustBeSetError(resp, "client_id")
		return
	}
	if config.PrivateKey.Value == "" {
		addAttributeMustBeSetError(resp, "private_key")
		return
	}
	if config.ApiVersion.Value == "" {
		addAttributeMustBeSetError(resp, "api_version")
		return
	}
	if config.Username.Value == "" {
		addAttributeMustBeSetError(resp, "username")
		return
	}
	client, err := auth.Client(auth.Config{
		ApiVersion: config.ApiVersion.Value,
		Username:   config.Username.Value,
		ClientId:   config.ClientId.Value,
		PrivateKey: config.PrivateKey.Value,
		LoginUrl:   config.LoginUrl.Value,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating salesforce client", err.Error())
		return
	}
	p.client = client
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"salesforce_profile":   profileType{},
		"salesforce_user":      userType{},
		"salesforce_user_role": userRoleType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"salesforce_profile":      profileDatasourceType{},
		"salesforce_user_license": userLicenseDatasourceType{},
	}, nil
}

func errorConvertingProvider(typ interface{}) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic("Error converting provider", fmt.Sprintf("An unexpected error was encountered converting the provider. This is always a bug in the provider.\n\nType: %T", typ))
}

func addAttributeMustBeSetError(resp *tfsdk.ConfigureProviderResponse, attr string) {
	resp.Diagnostics.AddAttributeError(
		tftypes.NewAttributePath().WithAttributeName(attr),
		"Invalid provider config",
		fmt.Sprintf("%s must be set.", attr),
	)
}

func addCannotInterpolateInProviderBlockError(resp *tfsdk.ConfigureProviderResponse, attr string) {
	resp.Diagnostics.AddAttributeError(
		tftypes.NewAttributePath().WithAttributeName(attr),
		"Can't interpolate into provider block",
		"Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
	)
}
