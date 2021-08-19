package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var providerFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"salesforce": func() (tfprotov6.ProviderServer, error) {
		return tfsdk.NewProtocol6Server(New()), nil
	},
}

func testAccPreCheck(t *testing.T) {
}
