package common

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Salesforce displays some IDs in the old 15 digit case sensitive format (such as in the url)
// if a user pastes the old format in their config it leads to permadiffs with the ids read from the API
// the conversion source code is posted here
// https://help.salesforce.com/articleView?id=000319308&type=1&mode=1
func NormalizeId(id string) string {
	if len(id) != 15 {
		// if the string is empty or already 18 characters, or not a proper id, just return it
		// let it error somewhere else
		return id
	}
	var addon string
	for block := 0; block < 3; block++ {
		loop := 0
		for position := 0; position < 5; position++ {
			current := id[block*5+position]
			if current >= 'A' && current <= 'Z' {
				loop += 1 << position
			}
		}
		addon += string("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"[loop])
	}
	return id + addon
}

func SuppressIdDiff(k, old, new string, d *schema.ResourceData) bool {
	return NormalizeId(old) == NormalizeId(new)
}
