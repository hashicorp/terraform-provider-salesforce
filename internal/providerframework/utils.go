package providerframework

import "github.com/hashicorp/terraform-plugin-go/tfprotov6"

func errToDiags(err error) []*tfprotov6.Diagnostic {
	return []*tfprotov6.Diagnostic{
		{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  err.Error(),
		},
	}
}
