data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

output "profile_id" {
  value = data.salesforce_profile.chatter_free.id
}