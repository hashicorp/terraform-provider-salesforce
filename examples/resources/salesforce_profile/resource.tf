data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

resource "salesforce_profile" "example" {
  name            = "example"
  user_license_id = data.salesforce_user_license.fdc.id
  description     = "example"
}