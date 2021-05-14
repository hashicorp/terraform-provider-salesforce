package providerdynamic

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-salesforce/internal/common"
)

// ConfigureProvider function
func (s *providerServer) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	response := &tfprotov5.ConfigureProviderResponse{}
	var providerConfig map[string]tftypes.Value
	var err error

	// transform provider config schema into tftype.Type and unmarshal the given config into a tftypes.Value
	cfgType := GetTypeFromSchema(GetProviderConfigSchema())
	cfgVal, err := req.Config.Unmarshal(cfgType)
	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to decode ConfigureProvider request parameter",
			Detail:   err.Error(),
		})
		return response, nil
	}

	err = cfgVal.As(&providerConfig)
	if err != nil {
		// invalid configuration schema - this shouldn't happen, bail out now
		response.Diagnostics = append(response.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Provider configuration: failed to extract values",
			Detail:   err.Error(),
		})
		return response, nil
	}

	client, err := common.Client(common.Config{
		ClientId:   EnvIfNull(providerConfig["client_id"], "SALESFORCE_CLIENT_ID"),
		PrivateKey: EnvIfNull(providerConfig["private_key"], "SALESFORCE_PRIVATE_KEY_FILE", "SALESFORCE_PRIVATE_KEY"),
		ApiVersion: EnvIfNull(providerConfig["api_version"], "SALESFORCE_API_VERSION"),
		Username:   EnvIfNull(providerConfig["api_version"], "SALESFORCE_USERNAME"),
	})
	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov5.Diagnostic{
			Severity: tfprotov5.DiagnosticSeverityError,
			Summary:  "Unable to create salesforce client",
			Detail:   err.Error(),
		})
		return response, nil
	}

	s.client = client
	return response, nil
}

func EnvIfNull(val tftypes.Value, env ...string) string {
	if val.IsNull() {
		for _, e := range env {
			if v := os.Getenv(e); v != "" {
				return v
			}
		}
	}
	var v string
	val.As(&v)
	return v
}
