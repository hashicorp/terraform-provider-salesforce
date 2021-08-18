package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/nimajalali/go-force/force"
)

type userType struct {
}

func (u userType) GetSchema(_ context.Context) (tfsdk.Schema, []*tfprotov6.Diagnostic) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"alias": {
				Type:     types.StringType,
				Required: true,
			},
			"email": {
				Type:     types.StringType,
				Required: true,
			},
			"email_encoding_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "ISO-8859-1",
			},
			"is_active": {
				Type:     types.BoolType,
				Optional: true,
				// Default:  true,
			},
			"language_locale_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "en_US",
			},
			"last_name": {
				Type:     types.StringType,
				Required: true,
			},
			"locale_sid_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "en_US",
			},
			"profile_id": {
				Type:     types.StringType,
				Required: true,
			},
			"time_zone_sid_key": {
				Type:     types.StringType,
				Optional: true,
				// Default:  "America/Los_Angeles",
			},
			"username": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (u userType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	prov, ok := p.(*provider)
	if !ok {
		return nil, []*tfprotov6.Diagnostic{
			{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error converting provider",
				Detail:   fmt.Sprintf("An unexpected error was encountered converting the provider. This is always a bug in the provider.\n\nType: %T", p),
			},
		}
	}
	return userResource{client: prov.client}, nil
}

type userResource struct {
	client *force.ForceApi
}

type User struct {
	Id                types.String `tfsdk:"id" force:"-"`
	Alias             string       `tfsdk:"alias" force:",omitempty"`
	Email             string       `tfsdk:"email" force:",omitempty"`
	EmailEncodingKey  *string      `tfsdk:"email_encoding_key" force:",omitempty"`
	IsActive          *bool        `tfsdk:"is_active" force:",omitempty"`
	LanguageLocaleKey *string      `tfsdk:"language_locale_key" force:",omitempty"`
	LastName          string       `tfsdk:"last_name" force:",omitempty"`
	LocaleSidKey      *string      `tfsdk:"locale_sid_key" force:",omitempty"`
	ProfileID         string       `tfsdk:"profile_id" force:",omitempty"`
	TimeZoneSidKey    *string      `tfsdk:"time_zone_sid_key" force:",omitempty"`
	Username          string       `tfsdk:"username" force:",omitempty"`
}

func (u User) ApiName() string {
	return "User"
}

func (u User) ExternalIdApiName() string {
	return ""
}

func (u User) withDefaults() User {
	emailEncodingKey := "ISO-8859-1"
	isActive := true
	languageLocaleKey := "en_US"
	localeSidKey := "en_US"
	timeZoneSidKey := "America/Los_Angeles"

	if u.EmailEncodingKey == nil {
		u.EmailEncodingKey = &emailEncodingKey
	}
	if u.IsActive == nil {
		u.IsActive = &isActive
	}
	if u.LanguageLocaleKey == nil {
		u.LanguageLocaleKey = &languageLocaleKey
	}
	if u.LocaleSidKey == nil {
		u.LocaleSidKey = &localeSidKey
	}
	if u.TimeZoneSidKey == nil {
		u.TimeZoneSidKey = &timeZoneSidKey
	}

	return u
}

func (u userResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var user User
	if diags := req.Plan.Get(ctx, &user); diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}
	// non-pointer method will not mutate the user declared in this function
	// this is useful so that State.Set doesn't error with a mismatch with the plan
	// where the unset fields were of course, unset.
	sfResp, err := u.client.InsertSObject(user.withDefaults())
	if err != nil {
		resp.AddError("Error inserting User", err.Error())
		return
	}
	user.Id = types.String{Value: sfResp.Id}

	resp.Diagnostics = resp.State.Set(ctx, &user)
}

func (u userResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var user User
	if diags := req.State.Get(ctx, &user); diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}

	err := u.client.GetSObject(user.Id.Value, nil, &user)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
		} else {
			resp.AddError("Error getting User", err.Error())
		}
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &user)
}

func (u userResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var user User
	if diags := req.Plan.Get(ctx, &user); diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}
	// non-pointer method will not mutate the user declared in this function
	// this is useful so that State.Set doesn't error with a mismatch with the plan
	// where the unset fields were of course, unset.
	if err := u.client.UpdateSObject(user.Id.Value, user.withDefaults()); err != nil {
		resp.AddError("Error updating User", err.Error())
		return
	}

	resp.Diagnostics = resp.State.Set(ctx, &user)
}

func (u userResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	id, diags := req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"))
	if diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}

	isActive := false
	err := u.client.UpdateSObject(id.(types.String).Value, User{IsActive: &isActive})
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
		} else {
			resp.AddError("Error deleting User", err.Error())
		}
		return
	}

	resp.State.RemoveResource(ctx)
	resp.AddWarning("Users cannot be deleted from salesforce", "Destroy has deactivated the user and discarded it from Terraform state, but the record continues to exist, and the unique username remains taken")
}
