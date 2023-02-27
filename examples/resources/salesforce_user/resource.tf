# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

resource "salesforce_user_role" "ceo" {
  name           = "ceo"
  developer_name = "ceo"
}

resource "salesforce_user" "example" {
  alias               = "example"
  email               = "user@example.com"
  last_name           = "example"
  username            = "user@example.com"
  profile_id          = data.salesforce_profile.chatter_free.id
  user_role_id        = salesforce_user_role.ceo.id
  email_encoding_key  = "UTF-8"
  language_locale_key = "en_US"
  time_zone_sid_key   = "America/Chicago"
  locale_sid_key      = "en_US"
  reset_password      = true
}