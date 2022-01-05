package fuzzconfig

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func FuzzSyntaxParseConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := hclsyntax.ParseConfig(data, "<fuzz-conf>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		return
	})
}

func FuzzSyntaxExpr(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := hclsyntax.ParseConfig(data, "<fuzz-conf>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		return
	})
}

