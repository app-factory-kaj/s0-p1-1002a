package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doGetHello(t *testing.T, target string) Greeting {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	GetHello(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	var g Greeting
	if err := json.NewDecoder(rec.Body).Decode(&g); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return g
}

func TestGetHelloWithName(t *testing.T) {
	g := doGetHello(t, "/hello?name=Ada")
	if g.Name != "Ada" {
		t.Errorf("name = %q, want %q", g.Name, "Ada")
	}
	if g.Message == "" {
		t.Error("message is empty")
	}
}

func TestGetHelloDefaultsWhenMissing(t *testing.T) {
	g := doGetHello(t, "/hello")
	if g.Name != "World" {
		t.Errorf("name = %q, want %q", g.Name, "World")
	}
}

func TestGetHelloDefaultsWhenEmpty(t *testing.T) {
	g := doGetHello(t, "/hello?name=")
	if g.Name != "World" {
		t.Errorf("name = %q, want %q", g.Name, "World")
	}
}
