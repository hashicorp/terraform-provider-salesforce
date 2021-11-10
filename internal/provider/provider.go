package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"

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
				Description: "API version of the salesforce org in the format in the format: vMAJOR.MINOR. The provider requires at least version v53.0. Can be specified with the environment variable SALESFORCE_API_VERSION.",
				Type:        types.StringType,
				Optional:    true,
			},
			"username": {
				Description: "Salesforce Username of a System Administrator like user for the provider to authenticate as. Can be specified with the environment variable SALESFORCE_USERNAME.",
				Type:        types.StringType,
				Optional:    true,
			},
			"is_sandbox_org": {
				Description: "Indicates if the salesforce org is a sandbox org or a developer/production org. Ensures the provider attempts to authenticate with the correct server. Can be specified with the environment variable SALESFORCE_IS_SANDBOX_ORG.",
				Type:        types.BoolType,
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
	Sandbox    types.Bool   `tfsdk:"is_sandbox_org"`
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
	if config.Sandbox.Unknown {
		addCannotInterpolateInProviderBlockError(resp, "is_sandbox_org")
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
	if config.Sandbox.Null {
		if isSandboxStr := os.Getenv("SALESFORCE_IS_SANDBOX_ORG"); isSandboxStr != "" {
			isSandbox, err := strconv.ParseBool(isSandboxStr)
			if err != nil {
				resp.Diagnostics.AddAttributeError(
					tftypes.NewAttributePath().WithAttributeName("is_sandbox_org"),
					"Invalid provider config",
					err.Error(),
				)
				return
			}
			config.Sandbox.Value = isSandbox
		}
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
		Sandbox:    config.Sandbox.Value,
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
