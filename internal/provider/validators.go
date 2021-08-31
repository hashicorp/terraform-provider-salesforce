package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type descriptions struct {
}

func (descriptions) Description(ctx context.Context) string {
	return ""
}

func (descriptions) MarkdownDescription(ctx context.Context) string {
	return ""
}

type emptyString struct {
	descriptions
}

func (emptyString) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	if attr.Value == "" {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Attribute: req.AttributePath,
			Summary:   "Empty string",
		})
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
	descriptions
}

func (email) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	if !isEmailValid(attr.Value) {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Attribute: req.AttributePath,
			Summary:   "Invalid email address",
		})
	}
}

type stringInSlice struct {
	slice []string
	descriptions
}

func (s stringInSlice) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	attr := req.AttributeConfig.(types.String)
	if attr.Unknown {
		return
	}
	for _, item := range s.slice {
		if attr.Value == item {
			return
		}
	}
	resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
		Severity:  tfprotov6.DiagnosticSeverityError,
		Attribute: req.AttributePath,
		Summary:   "Invalid string",
		Detail:    fmt.Sprintf("String must be one of: [%s]", strings.Join(s.slice, ", ")),
	})
}
