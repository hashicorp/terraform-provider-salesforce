package provider

import "github.com/hashicorp/terraform-plugin-go/tfprotov6"

func errToDiags(err error) []*tfprotov6.Diagnostic {
	return []*tfprotov6.Diagnostic{
		{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  err.Error(),
		},
	}
}

func diagsHasError(diags []*tfprotov6.Diagnostic) bool {
	for _, diag := range diags {
		if diag == nil {
			continue
		}
		if diag.Severity == tfprotov6.DiagnosticSeverityError {
			return true
		}
	}

	return false
}
