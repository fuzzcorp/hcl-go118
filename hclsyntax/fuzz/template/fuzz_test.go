package fuzztemplate

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func FuzzSyntaxTemplate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := hclsyntax.ParseTemplate(data, "<fuzz-tmpl>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		return
	})
}
