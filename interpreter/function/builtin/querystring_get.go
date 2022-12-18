// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of querystring.get
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-get/
func Querystring_get(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Builtin function querystring.get is not impelemented"))
}
