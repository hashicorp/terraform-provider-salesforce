package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
)

type Profile struct {
	sobjects.BaseSObject
}

type ProfileQueryResponse struct {
	sobjects.BaseQuery
	Records []Profile `json:"Records" force:"records"`
}

func datasourceProfile() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ReadContext: datasourceProfileRead,
	}
}

func datasourceProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var query ProfileQueryResponse
	nameFilter := fmt.Sprintf("Name = '%s'", d.Get("name").(string))
	err := meta.(*force.ForceApi).Query(force.BuildQuery("Id, Name", "Profile", []string{nameFilter}), &query)
	if err != nil {
		return diag.FromErr(err)
	}
	profile := query.Records[0]
	d.SetId(profile.Id)
	d.Set("name", profile.Name)
	return nil
}
