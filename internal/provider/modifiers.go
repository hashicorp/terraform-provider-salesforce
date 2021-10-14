package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Salesforce displays some IDs in the old 15 digit case sensitive format (such as in the url)
// if a user pastes the old format in their config it leads to permadiffs with the ids read from the API
// the conversion source code is posted here
// https://help.salesforce.com/articleView?id=000319308&type=1&mode=1
func normalizeId(id string) string {
	if len(id) != 15 {
		// if the string is empty or already 18 characters, or not a proper id, just return it
		// let it error upstream
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

type NormalizeId struct {
	emptyDescriptions
}

func (NormalizeId) Modify(_ context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if req.AttributeState == nil {
		resp.AttributePlan = req.AttributePlan
		return
	}
	plan := req.AttributePlan.(types.String)
	state := req.AttributeState.(types.String)
	if normalizeId(plan.Value) == normalizeId(state.Value) {
		resp.AttributePlan = state
	} else {
		resp.AttributePlan = plan
	}
}

type resourceDefaults struct {
	defaults map[string]attr.Value
	emptyDescriptions
}

func (r resourceDefaults) Modify(_ context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	resp.AttributePlan = req.AttributePlan
	// TODO need to check if null, not supported generically for attr.Value
	// for now all defaults happen to be string so this is fine for now
	if req.AttributeConfig.(types.String).Null {
		def, ok := r.defaults[req.AttributePath.String()]
		if ok {
			resp.AttributePlan = def
		}
	}
}

type staticComputed struct {
	emptyDescriptions
}

func (staticComputed) Modify(_ context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if req.AttributeState == nil {
		return
	}
	resp.AttributePlan = req.AttributeState
}

type optionalComputed struct {
	emptyDescriptions
}

func (optionalComputed) Modify(_ context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if req.AttributeState == nil {
		return
	}
	// TODO this is bool specific, ideally we could check unset more generically
	if req.AttributeConfig == nil || req.AttributeConfig.Equal(types.Bool{Null: true}) {
		resp.AttributePlan = req.AttributeState
	}
}
