package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProfile_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProfile_basic("Chatter Free User"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.salesforce_profile.test", "id"),
				),
			},
		},
	})
}

func testAccDataSourceProfile_basic(name string) string {
	return fmt.Sprintf(`
data "salesforce_profile" "test" {
  name = "%s"
}
`, name)
}
