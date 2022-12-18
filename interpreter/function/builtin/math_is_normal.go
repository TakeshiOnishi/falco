// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of math.is_normal
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-normal/
func Math_is_normal(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Builtin function math.is_normal is not impelemented"))
}
