package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUserRole_basic(t *testing.T) {
	t.Parallel()

	developerName := fmt.Sprintf("tf_test_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserRole_basic(developerName),
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

	developerName := fmt.Sprintf("tf_test_%s", acctest.RandString(10))
	developerNameParent := fmt.Sprintf("tf_test_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserRole_basic(developerName),
			},
			{
				ResourceName:      "salesforce_user_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUserRole_with_parent(developerNameParent, developerName),
			},
			{
				ResourceName:      "salesforce_user_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUserRole_with_parent_no_assign(developerNameParent, developerName),
			},
		},
	})
}

func testAccResourceUserRole_basic(developerName string) string {
	return fmt.Sprintf(`
resource "salesforce_user_role" "test" {
  name           = "test"
  developer_name = "%s"
}
`, developerName)
}

func testAccResourceUserRole_with_parent(developerNameParent, developerName string) string {
	return fmt.Sprintf(`
resource "salesforce_user_role" "parent" {
  name           = "parent"
  developer_name = "%s"
}

resource "salesforce_user_role" "test" {
  name           = "child"
  developer_name = "%s"
  parent_role_id = resource.salesforce_user_role.parent.id
}
`, developerNameParent, developerName)
}

func testAccResourceUserRole_with_parent_no_assign(developerNameParent, developerName string) string {
	return fmt.Sprintf(`
resource "salesforce_user_role" "parent" {
  name           = "parent"
  developer_name = "%s"
}

resource "salesforce_user_role" "test" {
  name           = "child"
  developer_name = "%s"
}
`, developerNameParent, developerName)
}
