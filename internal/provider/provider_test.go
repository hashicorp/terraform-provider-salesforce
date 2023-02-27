// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var providerFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"salesforce": func() (tfprotov6.ProviderServer, error) {
		return providerserver.NewProtocol6WithError(New())()
	},
}

func testAccPreCheck(t *testing.T) {
	testEnvVars := []string{"SALESFORCE_CLIENT_ID", "SALESFORCE_PRIVATE_KEY", "SALESFORCE_API_VERSION", "SALESFORCE_USERNAME"}
	for _, env := range testEnvVars {
		if os.Getenv(env) == "" {
			t.Fatalf("%s must be set for acceptance tests", env)
		}
	}
}
