package curlize

import (
	"net/http"
	"strings"
	"testing"
)

func TestCurlize_Get(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "http://localhost:80/foo/bar", strings.NewReader("foo/bar"))
	cmd, err := Curlize(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := `curl -X GET -d foo/bar http://localhost:80/foo/bar`
	if cmd.String() != expected {
		t.Fatalf("expected command is %+v, got %+v", expected, cmd.String())
	}
}

func TestCurlize_GetWithEmptyBody(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "http://localhost:80/foo/bar", nil)
	cmd, err := Curlize(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := `curl -X GET http://localhost:80/foo/bar`
	if cmd.String() != expected {
		t.Fatalf("expected command is %+v, got %+v", expected, cmd.String())
	}
}

func TestCurlize_Post(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "http://localhost:80/foo/bar", strings.NewReader(`{"foo": "bar"}`))
	r.Header.Set("Content-Type", "application/json")
	cmd, err := Curlize(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := `curl -X POST -d '{"foo": "bar"}' -H 'Content-Type: application/json' http://localhost:80/foo/bar`
	if cmd.String() != expected {
		t.Fatalf("expected command is %+v, got %+v", expected, cmd.String())
	}
}
