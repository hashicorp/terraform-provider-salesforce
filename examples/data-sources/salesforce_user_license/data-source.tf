# for list of keys see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_userlicense.htm
data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

output "user_license_id" {
  value = data.salesforce_user_license.fdc.id
}