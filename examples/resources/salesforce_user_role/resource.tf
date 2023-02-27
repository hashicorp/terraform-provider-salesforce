# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "salesforce_user_role" "parent" {
  name           = "parent"
  developer_name = "parent"
}

resource "salesforce_user_role" "child" {
  name           = "child"
  developer_name = "child"
  parent_role_id = resource.salesforce_user_role.parent.id
}