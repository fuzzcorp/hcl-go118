package fuzztraversal

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func FuzzSyntaxTraversal(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := hclsyntax.ParseTraversalAbs(data, "<fuzz-trav>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		return
	})
}
