package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-salesforce/internal/common"
	"github.com/nimajalali/go-force/force"
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	client *force.ForceApi
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, []*tfprotov6.Diagnostic) {
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
	if diags := req.Config.Get(ctx, &config); diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}

	// interpolation not allowed in provider block
	if config.ClientId.Unknown {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Can't interpolate into provider block",
			Detail:    "Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("client_id"),
		})
		return
	}
	if config.PrivateKey.Unknown {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Can't interpolate into provider block",
			Detail:    "Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("private_key"),
		})
		return
	}
	if config.ApiVersion.Unknown {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Can't interpolate into provider block",
			Detail:    "Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("api_version"),
		})
		return
	}
	if config.Username.Unknown {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Can't interpolate into provider block",
			Detail:    "Interpolating that value into the provider block doesn't give the provider enough information to run. Try hard-coding the value, instead.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("username"),
		})
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
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Invalid provider config",
			Detail:    "client_id must be set.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("client_id"),
		})
		return
	}
	if config.PrivateKey.Value == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Invalid provider config",
			Detail:    "private_key must be set.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("private_key"),
		})
		return
	}
	if config.ApiVersion.Value == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Invalid provider config",
			Detail:    "api_version must be set.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("api_version"),
		})
		return
	}
	if config.Username.Value == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Invalid provider config",
			Detail:    "username must be set.",
			Attribute: tftypes.NewAttributePath().WithAttributeName("username"),
		})
		return
	}
	client, err := common.Client(common.Config{
		ApiVersion: config.ApiVersion.Value,
		Username:   config.Username.Value,
		ClientId:   config.ClientId.Value,
		PrivateKey: config.PrivateKey.Value,
	})
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error creating salesforce client",
			Detail:   err.Error(),
		})
		return
	}
	p.client = client
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.ResourceType{
		"salesforce_user": userType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	return map[string]tfsdk.DataSourceType{
		"salesforce_profile": profileType{},
	}, nil
}
