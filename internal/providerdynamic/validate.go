package providerdynamic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// ValidateResourceTypeConfig function
func (s *providerServer) ValidateResourceTypeConfig(ctx context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	return &tfprotov5.ValidateResourceTypeConfigResponse{}, nil
}
