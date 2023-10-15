package interpreter

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
	"github.com/ysugimoto/falco/resolver"
)

func defaultBackend(url *url.URL) string {
	return fmt.Sprintf(`
backend example {
  .host = "%s";
  .port = "%s";
  .ssl = false;
}
`, url.Hostname(), url.Port(),
	)
}

func assertInterpreter(t *testing.T, vcl string, scope context.Scope, assertions map[string]value.Value) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	// server.EnableHTTP2 = true
	defer server.Close()

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Errorf("Test server URL parsing error: %s", err)
		return
	}

	vcl = defaultBackend(parsed) + "\n" + vcl
	ip := New(context.WithResolver(
		resolver.NewStaticResolver("main", vcl),
	))
	ip.ServeHTTP(
		httptest.NewRecorder(),
		httptest.NewRequest(http.MethodGet, "http://localhost", nil),
	)

	for name, val := range assertions {
		v, err := ip.vars.Get(scope, name)
		if err != nil {
			t.Errorf("Value get error: %s", err)
			return
		} else if v == nil || v == value.Null {
			t.Errorf("Value %s is nil", name)
			return
		}
		if diff := cmp.Diff(val, v); diff != "" {
			t.Errorf("Value asserion error, diff: %s", diff)
		}
	}
}

func assertValue(t *testing.T, name string, expect, actual value.Value) {
	if expect.Type() != actual.Type() {
		t.Errorf("%s type unmatch, expect %s, got %s", name, expect.Type(), actual.Type())
		return
	}
	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Errorf("Value asserion error, diff: %s", diff)
	}
}