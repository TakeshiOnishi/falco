// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of std.itoa_charset
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa-charset/
func Std_itoa_charset(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Not Impelemented"))
}
