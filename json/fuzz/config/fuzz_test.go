package fuzzconfig

import (
	"testing"

	"github.com/hashicorp/hcl/v2/json"
)

func FuzzJSONConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := json.Parse(data, "<fuzz-conf>")

		if diags.HasErrors() {
			return
		}

		return
	})
}
