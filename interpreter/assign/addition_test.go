package assign

import (
	"net"
	"testing"
	"time"

	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/interpreter/value"
)

func TestProcessAddition(t *testing.T) {
	t.Run("left is INTEGER", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    int64
			right   value.Value
			expect  int64
			isError bool
		}{
			{left: 10, right: &value.Integer{Value: 100}, expect: 110},
			{left: 10, right: &value.Integer{Value: 100, Literal: true}, expect: 110},
			{left: 10, right: &value.Float{Value: 50.0}, expect: 60},
			{left: 10, right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: 10, right: &value.String{Value: "example"}, isError: true},
			{left: 10, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: 10, right: &value.RTime{Value: 100 * time.Second}, expect: 110},
			{left: 10, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: 10, right: &value.Time{Value: now}, expect: 10 + now.Unix()},
			{left: 10, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: 10, right: &value.Boolean{Value: true}, isError: true},
			{left: 10, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: 10, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.Integer{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %d, got %d", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is FLOAT", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    float64
			right   value.Value
			expect  float64
			isError bool
		}{
			{left: 10.0, right: &value.Integer{Value: 100}, expect: 110.0},
			{left: 10.0, right: &value.Integer{Value: 100, Literal: true}, expect: 110.0},
			{left: 10, right: &value.Float{Value: 50.0}, expect: 60.0},
			{left: 10, right: &value.Float{Value: 50.0, Literal: true}, expect: 60.0},
			{left: 10, right: &value.String{Value: "example"}, isError: true},
			{left: 10, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: 10, right: &value.RTime{Value: 100 * time.Second}, expect: 10 + float64(100)},
			{left: 10, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: 10, right: &value.Time{Value: now}, expect: 10 + float64(now.Unix())},
			{left: 10, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: 10, right: &value.Boolean{Value: true}, isError: true},
			{left: 10, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: 10, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.Float{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %.2f, got %.2f", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is STRING", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    string
			right   value.Value
			expect  string
			isError bool
		}{
			{left: "left", right: &value.Integer{Value: 100}, isError: true},
			{left: "left", right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: "left", right: &value.Float{Value: 50.0}, isError: true},
			{left: "left", right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: "left", right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: "left", right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: "left", right: &value.Time{Value: now}, isError: true},
			{left: "left", right: &value.String{Value: "example"}, isError: true},
			{left: "left", right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: "left", right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: "left", right: &value.Boolean{Value: true}, isError: true},
			{left: "left", right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: "left", right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.String{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is RTIME", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    time.Duration
			right   value.Value
			expect  time.Duration
			isError bool
		}{
			{left: time.Second, right: &value.Integer{Value: 100}, expect: 101 * time.Second},
			{left: time.Second, right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: time.Second, right: &value.Float{Value: 50.0}, expect: 51 * time.Second},
			{left: time.Second, right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: time.Second, right: &value.String{Value: "example"}, isError: true},
			{left: time.Second, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: time.Second, right: &value.RTime{Value: 100 * time.Second}, expect: 101 * time.Second},
			{left: time.Second, right: &value.RTime{Value: 100 * time.Second, Literal: true}, expect: 101 * time.Second},
			{left: time.Second, right: &value.Time{Value: now}, expect: time.Second + time.Duration(now.Unix())},
			{left: time.Second, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: time.Second, right: &value.Boolean{Value: true}, isError: true},
			{left: time.Second, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: time.Second, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.RTime{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is TIME", func(t *testing.T) {
		now := time.Now()
		now2 := now.Add(10 * time.Second)
		tests := []struct {
			left    time.Time
			right   value.Value
			expect  time.Time
			isError bool
		}{
			{left: now, right: &value.Integer{Value: 100}, expect: now.Add(100 * time.Second)},
			{left: now, right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: now, right: &value.Float{Value: 50.0}, expect: now.Add(50 * time.Second)},
			{left: now, right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: now, right: &value.String{Value: "example"}, isError: true},
			{left: now, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: now, right: &value.RTime{Value: 100 * time.Second}, expect: now.Add(100 * time.Second)},
			{left: now, right: &value.RTime{Value: 100 * time.Second, Literal: true}, expect: now.Add(100 * time.Second)},
			{left: now, right: &value.Time{Value: now2}, isError: true},
			{left: now, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: now, right: &value.Boolean{Value: true}, isError: true},
			{left: now, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: now, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.Time{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is BACKEND", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    string
			right   value.Value
			expect  string
			isError bool
		}{
			{left: "backend", right: &value.Integer{Value: 100}, isError: true},
			{left: "backend", right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: "backend", right: &value.Float{Value: 50.0}, isError: true},
			{left: "backend", right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: "backend", right: &value.String{Value: "example"}, isError: true},
			{left: "backend", right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: "backend", right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: "backend", right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: "backend", right: &value.Time{Value: now}, isError: true},
			{left: "backend", right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: "backend", right: &value.Boolean{Value: true}, isError: true},
			{left: "backend", right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: "backend", right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: tt.left}}}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value.Name.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value.Name.Value)
			}
		}
	})

	t.Run("left is BOOL", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    bool
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: false, right: &value.Integer{Value: 100}, isError: true},
			{left: false, right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: false, right: &value.Float{Value: 50.0}, isError: true},
			{left: false, right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: false, right: &value.String{Value: "example"}, isError: true},
			{left: false, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: false, right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: false, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: false, right: &value.Time{Value: now}, isError: true},
			{left: false, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: false, right: &value.Boolean{Value: true}, isError: true},
			{left: false, right: &value.Boolean{Value: true, Literal: true}, isError: true},
			{left: false, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			left := &value.Boolean{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is IP", func(t *testing.T) {
		now := time.Now()
		v := net.ParseIP("127.0.0.1")
		vv := net.ParseIP("127.0.0.2")
		tests := []struct {
			left    net.IP
			right   value.Value
			expect  net.IP
			isError bool
		}{
			{left: v, right: &value.Integer{Value: 100}, isError: true},
			{left: v, right: &value.Integer{Value: 100, Literal: true}, isError: true},
			{left: v, right: &value.Float{Value: 50.0}, isError: true},
			{left: v, right: &value.Float{Value: 50.0, Literal: true}, isError: true},
			{left: v, right: &value.String{Value: "example"}, isError: true},
			{left: v, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: v, right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: v, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: v, right: &value.Time{Value: now}, isError: true},
			{left: v, right: &value.String{Value: "127.0.0.2", Literal: true}, isError: true},
			{left: v, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: v, right: &value.Boolean{Value: true}, isError: true},
			{left: v, right: &value.Boolean{Value: true, Literal: true}, isError: true},
			{left: v, right: &value.IP{Value: vv}, isError: true},
		}

		for i, tt := range tests {
			left := &value.IP{Value: tt.left}
			err := Addition(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value.String() != tt.expect.String() {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})
}
