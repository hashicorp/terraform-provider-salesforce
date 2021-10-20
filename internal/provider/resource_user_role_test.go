package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUserRole_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserRole_basic(),
			},
			{
				ResourceName:      "salesforce_user_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUserRole_update(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserRole_basic(),
			},
			{
				ResourceName:      "salesforce_user_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUserRole_with_parent(),
			},
			{
				ResourceName:      "salesforce_user_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceUserRole_basic() string {
	return fmt.Sprintf(`
resource "salesforce_user_role" "test" {
  name           = "test"
  developer_name = "test"
}
`)
}

func testAccResourceUserRole_with_parent() string {
	return fmt.Sprintf(`
resource "salesforce_user_role" "parent" {
  name           = "parent"
  developer_name = "parent"
}

resource "salesforce_user_role" "test" {
  name           = "child"
  developer_name = "child"
  parent_role_id = resource.salesforce_user_role.parent.id
}
`)
}
