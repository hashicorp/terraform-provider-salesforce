package providerdynamic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// GetProviderSchema function
func (s *providerServer) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {

	cfgSchema := GetProviderConfigSchema()

	resSchema := GetProviderResourceSchema()

	return &tfprotov5.GetProviderSchemaResponse{
		Provider:        cfgSchema,
		ResourceSchemas: resSchema,
	}, nil
}
