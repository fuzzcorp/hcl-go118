package fuzzconfig

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func FuzzWriteParseConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		file, diags := hclwrite.ParseConfig(data, "<fuzz-conf>", hcl.Pos{Line: 1, Column: 1})

		if diags.HasErrors() {
			return
		}

		_, err := file.WriteTo(ioutil.Discard)

		if err != nil {
			return
		}

	})
}
