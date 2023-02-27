// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProfile_basic(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(name),
				// TODO there is a bug around maps that causes a permadiff for empty maps
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceProfile_update(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(name),
				// TODO there is a bug around maps that causes a permadiff for empty maps
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// importing permissions map not currently possible
				Config: testAccResourceProfile_with_permissions(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("salesforce_profile.test", "name", name),
					resource.TestCheckResourceAttr("salesforce_profile.test", "description", "test update"),
					resource.TestCheckResourceAttr("salesforce_profile.test", "permissions.%", "2"),
					resource.TestCheckResourceAttr("salesforce_profile.test", "permissions.EditTask", "true"),
					resource.TestCheckResourceAttr("salesforce_profile.test", "permissions.EmailSingle", "true"),
				),
			},
		},
	})
}

func testAccResourceProfile_basic(name string) string {
	return fmt.Sprintf(`
data "salesforce_user_license" "standard" {
  license_definition_key = "AUL"
}

resource "salesforce_profile" "test" {
  name            = "%s"
  user_license_id = data.salesforce_user_license.standard.id
  description     = "test"
}
`, name)
}

func testAccResourceProfile_with_permissions(name string) string {
	return fmt.Sprintf(`
data "salesforce_user_license" "standard" {
  license_definition_key = "AUL"
}

resource "salesforce_profile" "test" {
  name            = "%s"
  user_license_id = data.salesforce_user_license.standard.id
  description     = "test update"
  permissions = {
    EmailSingle = true
    EditTask = true
  }
}
`, name)
}
