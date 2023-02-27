# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "salesforce" {
  client_id   = "ABCDEFG"
  private_key = "/Users/mscott/priv.pem"
  api_version = "53.0"
  username    = "user@example.com"
}