package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-salesforce/internal/picklists"
	"github.com/nimajalali/go-force/force"
)

var userDefaults = resourceDefaults{
	defaults: map[string]attr.Value{
		tftypes.NewAttributePath().WithAttributeName("email_encoding_key").String(): types.String{
			Value: "UTF-8",
		},
		tftypes.NewAttributePath().WithAttributeName("language_locale_key").String(): types.String{
			Value: "en_US",
		},
		tftypes.NewAttributePath().WithAttributeName("locale_sid_key").String(): types.String{
			Value: "en_US",
		},
		tftypes.NewAttributePath().WithAttributeName("time_zone_sid_key").String(): types.String{
			Value: "America/New_York",
		},
	},
}

type userType struct {
}

func (userType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "ID of the resource.",
				Type:        types.StringType,
				Computed:    true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					staticComputed{},
				},
			},
			"alias": {
				Description: "The user’s alias. For example, jsmith.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
				},
			},
			"email": {
				Description: "The user’s email address.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					email{},
				},
			},
			"email_encoding_key": {
				Description: "The email encoding for the user, such as ISO-8859-1 or UTF-8. Defaults to UTF-8.",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Validators: []tfsdk.AttributeValidator{
					stringInSlice{
						slice:    picklists.EmailEncodingKeys,
						optional: true,
					},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					userDefaults,
				},
			},
			"language_locale_key": {
				Description: "The user’s language. Defaults to en_US.",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Validators: []tfsdk.AttributeValidator{
					stringInSlice{
						slice:    picklists.LanguageLocaleKeys,
						optional: true,
					},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					userDefaults,
				},
			},
			"last_name": {
				Description: "The user’s last name.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
				},
			},
			"locale_sid_key": {
				Description: "The value of the field affects formatting and parsing of values, especially numeric values, in the user interface. It doesn’t affect the API. The field values are named according to the language, and the country if necessary, using two-letter ISO codes. The set of names is based on the ISO standard. You can also manually set a user’s locale in the user interface, and then use that value for inserting or updating other users via the API. Defaults to en_US.",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Validators: []tfsdk.AttributeValidator{
					stringInSlice{
						slice:    picklists.LocaleSidKeys,
						optional: true,
					},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					userDefaults,
				},
			},
			"profile_id": {
				Description: "ID of the user’s Profile. Use this value to cache metadata based on profile.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					notEmptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
				},
			},
			"time_zone_sid_key": {
				Description: "A User time zone affects the offset used when displaying or entering times in the user interface. But the API doesn’t use a User time zone when querying or setting values. Values for this field are named using region and key city, according to ISO standards. You can also manually set one User time zone in the user interface, and then use that value for creating or updating other User records via the API. Defaults to America/New_York.",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Validators: []tfsdk.AttributeValidator{
					stringInSlice{
						slice:    picklists.TimeZoneSidKeys,
						optional: true,
					},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					userDefaults,
				},
			},
			"username": {
				Description: "Contains the name that a user enters to log in to the API or the user interface. The value for this field must be in the form of an email address, using all lowercase characters. It must also be unique across all organizations. If you try to create or update a User with a duplicate value for this field, the operation is rejected. Each inserted User also counts as a license. Every organization has a maximum number of licenses. If you attempt to exceed the maximum number of licenses by inserting User records, the create request is rejected.",
				Type:        types.StringType,
				Required:    true,
				Validators: []tfsdk.AttributeValidator{
					email{},
				},
			},
			"user_role_id": {
				Description: "ID of the user’s UserRole.",
				Type:        types.StringType,
				Optional:    true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
					fixNullToUnknown{},
				},
			},
		},
	}, nil
}

func (u userType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	prov, ok := p.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(u)}
	}
	return &userResource{
		Resource: Resource{
			Client: prov.client,
			Data:   &userResourceData{},
		},
	}, nil
}

type userResource struct {
	Resource
}

func (u *userResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	id, diags := req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"))
	if diags.HasError() {
		resp.Diagnostics = diags
		return
	}

	isActive := false
	err := u.Client.UpdateSObject(id.(types.String).Value, userResourceData{IsActive: &isActive})
	if err != nil {
		if isNotFoundError(err) {
			resp.State.RemoveResource(ctx)
		} else {
			resp.Diagnostics.AddError("Error Deleting User", err.Error())
		}
		return
	}

	resp.State.RemoveResource(ctx)
	resp.Diagnostics.AddWarning("Users cannot be deleted from salesforce", "Destroy has deactivated the user and discarded it from Terraform state, but the record continues to exist, and the unique username remains taken")
}

type userResourceData struct {
	Alias             string       `tfsdk:"alias" force:",omitempty"`
	Email             string       `tfsdk:"email" force:",omitempty"`
	EmailEncodingKey  string       `tfsdk:"email_encoding_key" force:",omitempty"`
	LanguageLocaleKey string       `tfsdk:"language_locale_key" force:",omitempty"`
	LastName          string       `tfsdk:"last_name" force:",omitempty"`
	LocaleSidKey      string       `tfsdk:"locale_sid_key" force:",omitempty"`
	ProfileID         string       `tfsdk:"profile_id" force:",omitempty"`
	TimeZoneSidKey    string       `tfsdk:"time_zone_sid_key" force:",omitempty"`
	Username          string       `tfsdk:"username" force:",omitempty"`
	UserRoleId        *string      `tfsdk:"user_role_id"`
	IsActive          *bool        `tfsdk:"-" force:",omitempty"`
	Id                types.String `tfsdk:"id" force:"-"`
}

func (userResourceData) ApiName() string {
	return "User"
}

func (userResourceData) ExternalIdApiName() string {
	return ""
}

func (u *userResourceData) Instance() force.SObject {
	return u
}

func (u *userResourceData) Insertable() force.SObject {
	return *u
}

func (u *userResourceData) Updatable() force.SObject {
	return *u
}

func (u *userResourceData) GetId() string {
	return u.Id.Value
}

func (u *userResourceData) SetId(id string) {
	u.Id = types.String{Value: id}
}
