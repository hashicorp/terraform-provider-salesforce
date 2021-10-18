package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser_basic(t *testing.T) {
	t.Parallel()
	t.Skip("Users cannot be deleted and there are limited licenses, skipping, comment out this line to run locally")

	email := os.Getenv("SALESFORCE_USERNAME")
	parts := strings.Split(email, "@")
	var username string
	if !strings.Contains(parts[0], "+") {
		username = fmt.Sprintf("%s+%s@%s", parts[0], acctest.RandString(10), parts[1])
	} else {
		username = fmt.Sprintf("%s-%s@%s", parts[0], acctest.RandString(10), parts[1])
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser_basic(email, username),
			},
			{
				ResourceName:      "salesforce_user.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUser_update(t *testing.T) {
	t.Parallel()
	t.Skip("Users cannot be deleted and there are limited licenses, skipping, comment out this line to run locally")

	email := os.Getenv("SALESFORCE_USERNAME")
	parts := strings.Split(email, "@")
	var username string
	if !strings.Contains(parts[0], "+") {
		username = fmt.Sprintf("%s+%s@%s", parts[0], acctest.RandString(10), parts[1])
	} else {
		username = fmt.Sprintf("%s-%s@%s", parts[0], acctest.RandString(10), parts[1])
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser_basic(email, username),
			},
			{
				ResourceName:      "salesforce_user.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceUser_full(email, username),
			},
			{
				ResourceName:      "salesforce_user.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceUser_basic(email, username string) string {
	return fmt.Sprintf(`
data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

resource "salesforce_user" "test" {
  alias = "test"
  email = "%s"
  last_name = "test"
  username = "%s"
  profile_id = data.salesforce_profile.chatter_free.id
}
`, email, username)
}

func testAccResourceUser_full(email, username string) string {
	return fmt.Sprintf(`
data "salesforce_profile" "standard" {
  name = "Standard User"
}

resource "salesforce_user" "test" {
  alias = "test"
  email = "%s"
  last_name = "test"
  username = "%s"
  profile_id = data.salesforce_profile.standard.id
  email_encoding_key  = "ISO-8859-1"
  language_locale_key = "en_US"
  time_zone_sid_key   = "America/Chicago"
  locale_sid_key      = "en_US"
}
`, email, username)
}
