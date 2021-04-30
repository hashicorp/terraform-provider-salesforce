package provider

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nimajalali/go-force/force"
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
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"salesforce_user": resourceUser(),
		},
	}

	p.ConfigureContextFunc = configure

	return p
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiVersion := d.Get("api_version").(string)
	username := d.Get("username").(string)
	clientId := d.Get("client_id").(string)
	privateKey := d.Get("private_key").(string)

	// try to read private key as file
	privateKeyBytes, err := ioutil.ReadFile(privateKey)
	if os.IsNotExist(err) {
		// assume private key was passed directly
		privateKeyBytes = []byte(privateKey)
	} else if err != nil {
		return nil, diag.FromErr(err)
	}

	signedJwt, err := SignJWT(privateKeyBytes, username, clientId)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	resp, err := Authenticate(signedJwt)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	client, err := force.CreateWithAccessToken(apiVersion, clientId, resp.AccessToken, resp.InstanceUrl)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
