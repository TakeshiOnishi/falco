// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of std.prefixof
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-prefixof/
func Test_Std_prefixof(t *testing.T) {
	tests := []struct {
		input  string
		prefix string
		expect bool
	}{
		{input: "greenhouse", prefix: "green", expect: true},
		{input: "greanhouse", prefix: "green", expect: false},
	}

	for i, tt := range tests {
		ret, err := Std_prefixof(
			&context.Context{},
			&value.String{Value: tt.input},
			&value.String{Value: tt.prefix},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.BooleanType {
			t.Errorf("[%d] Unexpected return type, expect=BOOL, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.Boolean](ret)
		if diff := cmp.Diff(tt.expect, v.Value); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}