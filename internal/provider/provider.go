package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-salesforce/internal/common"
)

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALESFORCE_CLIENT_ID", nil),
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SALESFORCE_PRIVATE_KEY_FILE", "SALESFORCE_PRIVATE_KEY"}, nil),
			},
			"api_version": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALESFORCE_API_VERSION", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALESFORCE_USERNAME", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"salesforce_profile": datasourceProfile(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"salesforce_user": resourceUser(),
		},
	}

	p.ConfigureContextFunc = configure

	return p
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	client, err := common.Client(common.Config{
		ApiVersion: d.Get("api_version").(string),
		Username:   d.Get("username").(string),
		ClientId:   d.Get("client_id").(string),
		PrivateKey: d.Get("private_key").(string),
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return client, nil
}
