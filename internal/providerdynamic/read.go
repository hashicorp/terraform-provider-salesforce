package providerdynamic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// ReadResource function
func (s *providerServer) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	return &tfprotov5.ReadResourceResponse{}, nil
}
