// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type notEmptyString struct {
	emptyDescriptions
}

func (notEmptyString) Validate(_ context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	if attr.Value == "" {
		resp.Diagnostics.AddAttributeError(req.AttributePath, "Empty String", "")
	}
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

type email struct {
	emptyDescriptions
}

func (email) Validate(_ context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	if !isEmailValid(attr.Value) {
		resp.Diagnostics.AddAttributeError(req.AttributePath, "Invalid email address", "")
	}
}

type stringInSlice struct {
	slice    []string
	optional bool
	emptyDescriptions
}

func (s stringInSlice) Validate(_ context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	if s.optional && attr.Null {
		return
	}
	for _, item := range s.slice {
		if attr.Value == item {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(req.AttributePath, "Invalid string", fmt.Sprintf("String must be one of: [%s]", strings.Join(s.slice, ", ")))
}
