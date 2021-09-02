---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "salesforce_profile Data Source - terraform-provider-salesforce"
subcategory: ""
description: |-
  
---

# salesforce_profile (Data Source)



## Example Usage

```terraform
data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

output "profile_id" {
  value = data.salesforce_profile.chatter_free.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)

### Read-Only

- **id** (String) The ID of this resource.

