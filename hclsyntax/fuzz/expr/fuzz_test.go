package fuzzexpr

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func FuzzSyntaxExpr(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := hclsyntax.ParseExpression(data, "<fuzz-expr>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		return
	})
}
