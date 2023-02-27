// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUserLicense_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUserLicense_basic("AUL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.salesforce_user_license.test", "id"),
				),
			},
		},
	})
}

func testAccDataSourceUserLicense_basic(name string) string {
	return fmt.Sprintf(`
data "salesforce_user_license" "test" {
  license_definition_key = "%s"
}
`, name)
}
