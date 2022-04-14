package handler

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

func TestProxyHandler_ServeHTTP(t *testing.T) {

	fixture := `<html><head><title>Hello</title></head><body><h1>Приве!т</h1></body></html>`
	expected := `<html><head><title>Hello</title></head><body><h1>Приве™!т</h1></body></html>`
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fixture))
		if err != nil {
			t.Fatalf(err.Error())
		}
	}))
	defer backendServer.Close()

	backendURL, err := url.Parse(backendServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	handler := &ProxyHandler{
		SiteAddress: backendURL,
	}

	srv := httptest.NewServer(handler)
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(respBody)
	if actual != expected {
		t.Errorf("Actual: %s expected: %s", actual, expected)
	}
}
