package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProfile_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(),
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceProfile_with_permissions(),
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceProfile_basic() string {
	return fmt.Sprintf(`
data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

resource "salesforce_profile" "test" {
  name            = "test"
  user_license_id = data.salesforce_user_license.fdc.id
  description     = "test"
}
`)
}

func testAccResourceProfile_with_permissions() string {
	return fmt.Sprintf(`
data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

resource "salesforce_profile" "test" {
  name            = "test"
  user_license_id = data.salesforce_user_license.fdc.id
  description     = "test update"
  permissions_email_single = true
  permissions_edit_task = true
}
`)
}
