package providerframework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/nimajalali/go-force/force"
)

type userType struct {
}

func (u userType) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"alias": {
				Type:     types.StringType,
				Required: true,
			},
			"email": {
				Type:     types.StringType,
				Required: true,
			},
			"email_encoding_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "ISO-8859-1",
			},
			"is_active": {
				Type:     types.BoolType,
				Optional: true,
				// Default:  true,
			},
			"language_locale_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "en_US",
			},
			"last_name": {
				Type:     types.StringType,
				Required: true,
			},
			"locale_sid_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "en_US",
			},
			"profile_id": {
				Type:     types.StringType,
				Required: true,
			},
			"time_zone_sid_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "America/Los_Angeles",
			},
			"username": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (u userType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	prov, ok := p.(*provider)
	if !ok {
		return nil, []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error converting provider",
				Detail:   fmt.Sprintf("An unexpected error was encountered converting the provider. This is always a bug in the provider.\n\nType: %T", p),
			},
		}
	}
	return userResource{client: prov.client}, nil
}

type userResource struct {
	client *force.ForceApi
}

func (u userResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
}

func (u userResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
}

func (u userResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

func (u userResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}
