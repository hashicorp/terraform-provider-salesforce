package providerdynamic

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// GetProviderConfigSchema contains the definitions of all configuration attributes
func GetProviderConfigSchema() *tfprotov5.Schema {
	b := tfprotov5.SchemaBlock{
		Attributes: []*tfprotov5.SchemaAttribute{
			{
				Name:     "client_id",
				Type:     tftypes.String,
				Optional: true,
				Computed: true,
			},
			{
				Name:     "private_key",
				Type:     tftypes.String,
				Optional: true,
				Computed: true,
			},
			{
				Name:     "api_version",
				Type:     tftypes.String,
				Optional: true,
				Computed: true,
			},
			{
				Name:     "username",
				Type:     tftypes.String,
				Optional: true,
				Computed: true,
			},
		},
	}

	return &tfprotov5.Schema{
		Version: 1,
		Block:   &b,
	}
}

// GetTypeFromSchema returns the equivalent tftypes.Type representation of a given tfprotov5.Schema
func GetTypeFromSchema(s *tfprotov5.Schema) tftypes.Type {
	schemaTypeAttributes := map[string]tftypes.Type{}
	for _, att := range s.Block.Attributes {
		schemaTypeAttributes[att.Name] = att.Type
	}
	return tftypes.Object{
		AttributeTypes: schemaTypeAttributes,
	}
}
