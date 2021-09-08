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
		Attributes: map[string]tfsdk.Attribute{
			"client_id": {
				Type:     types.StringType,
				Optional: true,
			},
			"private_key": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
			"api_version": {
				Type:     types.StringType,
				Optional: true,
			},
			"username": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

type providerData struct {
	ClientId   types.String `tfsdk:"client_id"`
	PrivateKey types.String `tfsdk:"private_key"`
	ApiVersion types.String `tfsdk:"api_version"`
	Username   types.String `tfsdk:"username"`
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
	})
	if err != nil {
		resp.AddError("Error creating salesforce client", err.Error())
		return
	}
	p.client = client
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"salesforce_user": userType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"salesforce_profile": profileType{},
	}, nil
}

func addAttributeMustBeSetError(resp *tfsdk.ConfigureProviderResponse, attr string) {
	resp.AddAttributeError(
		tftypes.NewAttributePath().WithAttributeName(attr),
		"Invalid provider config",
		fmt.Sprintf("%s must be set.", attr),
	)
}

func addCannotInterpolateInProviderBlockError(resp *tfsdk.ConfigureProviderResponse, attr string) {
	resp.AddAttributeError(
		tftypes.NewAttributePath().WithAttributeName(attr),
		"Can't interpolate into provider block",
		"Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
	)
}
