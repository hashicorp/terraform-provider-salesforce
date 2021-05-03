package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nimajalali/go-force/force"
	"github.com/nimajalali/go-force/sobjects"
)

type User struct {
	sobjects.BaseSObject
	Alias             string `force:",omitempty"`
	Email             string `force:",omitempty"`
	EmailEncodingKey  string `force:",omitempty"`
	LanguageLocaleKey string `force:",omitempty"`
	LastName          string `force:",omitempty"`
	LocaleSidKey      string `force:",omitempty"`
	ProfileId         string `force:",omitempty"`
	TimeZoneSidKey    string `force:",omitempty"`
	Username          string `force:",omitempty"`
	IsActive          *bool  `force:",omitempty"`
}

func (t *User) ApiName() string {
	return "User"
}

type UserQueryResponse struct {
	sobjects.BaseQuery
	Records []User `json:"Records" force:"records"`
}

func UserFromResourceData(d *schema.ResourceData) *User {
	// go-force must use pointer type for booleans to discern between false and unset
	isActive := d.Get("is_active").(bool)
	return &User{
		Alias:             d.Get("alias").(string),
		Email:             d.Get("email").(string),
		EmailEncodingKey:  d.Get("email_encoding_key").(string),
		LanguageLocaleKey: d.Get("language_locale_key").(string),
		LastName:          d.Get("last_name").(string),
		LocaleSidKey:      d.Get("locale_sid_key").(string),
		ProfileId:         d.Get("profile_id").(string),
		TimeZoneSidKey:    d.Get("time_zone_sid_key").(string),
		Username:          d.Get("username").(string),
		IsActive:          &isActive,
	}
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_encoding_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ISO-8859-1",
			},
			"is_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"language_locale_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "en_US",
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"locale_sid_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "en_US",
			},
			"profile_id": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return normalizeId(old) == normalizeId(new)
				},
			},
			"time_zone_sid_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "America/Los_Angeles",
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resp, err := meta.(*force.ForceApi).InsertSObject(UserFromResourceData(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Id)
	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var user User
	err := meta.(*force.ForceApi).GetSObject(d.Id(), nil, &user)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(user.Id)
	d.Set("alias", user.Alias)
	d.Set("email", user.Email)
	d.Set("email_encoding_key", user.EmailEncodingKey)
	d.Set("is_active", user.IsActive)
	d.Set("language_locale_key", user.LanguageLocaleKey)
	d.Set("last_name", user.LastName)
	d.Set("locale_sid_key", user.LocaleSidKey)
	d.Set("profile_id", user.ProfileId)
	d.Set("time_zone_sid_key", user.TimeZoneSidKey)
	d.Set("username", user.Username)
	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := meta.(*force.ForceApi).UpdateSObject(d.Id(), UserFromResourceData(d))
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// users can only be deactivated
	isActive := false
	err := meta.(*force.ForceApi).UpdateSObject(d.Id(), &User{IsActive: &isActive})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diag.Diagnostics{
		{
			Severity: diag.Warning,
			Summary:  "Users cannot be deleted from salesforce",
			Detail:   "Destroy has deactivated the user and discarded it from Terraform state, but the record continues to exist, and the unique username remains taken",
		},
	}
}
