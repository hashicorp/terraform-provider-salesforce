# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

output "profile_id" {
  value = data.salesforce_profile.chatter_free.id
}