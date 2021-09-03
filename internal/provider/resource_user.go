package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
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
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"email": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					email{},
				},
			},
			"email_encoding_key": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
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
				Type:     types.StringType,
				Optional: true,
				Computed: true,
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
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"locale_sid_key": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
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
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
				},
			},
			"time_zone_sid_key": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
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
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					email{},
				},
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
	IsActive          *bool        `tfsdk:"-" force:",omitempty"`
	Alias             string       `tfsdk:"alias" force:",omitempty"`
	Email             string       `tfsdk:"email" force:",omitempty"`
	EmailEncodingKey  string       `tfsdk:"email_encoding_key" force:",omitempty"`
	LanguageLocaleKey string       `tfsdk:"language_locale_key" force:",omitempty"`
	LastName          string       `tfsdk:"last_name" force:",omitempty"`
	LocaleSidKey      string       `tfsdk:"locale_sid_key" force:",omitempty"`
	ProfileID         string       `tfsdk:"profile_id" force:",omitempty"`
	TimeZoneSidKey    string       `tfsdk:"time_zone_sid_key" force:",omitempty"`
	Username          string       `tfsdk:"username" force:",omitempty"`
}

func (u User) ApiName() string {
	return "User"
}

func (u User) ExternalIdApiName() string {
	return ""
}

func (u userResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var user User
	if diags := req.Plan.Get(ctx, &user); diagsHasError(diags) {
		resp.Diagnostics = diags
		return
	}

	sfResp, err := u.client.InsertSObject(user)
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

	if err := u.client.UpdateSObject(user.Id.Value, user); err != nil {
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
