package handler

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type HeadersChan chan map[string][]string
type BodyChan chan []byte

//func get(server, url, respChan ) {
//
//}

func TestProxyHandler_ServeHTTP(t *testing.T) {
	fixture := `<html><head><title>Hello</title></head><body><h1>Приве!т</h1></body></html>`
	expected := `<html><head><title>Hello</title></head><body><h1>Приве™!т</h1></body></html>`

	headersChan := make(HeadersChan)
	bodyChan := make(BodyChan)

	fixtureServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headersChan <- r.Header
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fixture))
		if err != nil {
			t.Fatalf(err.Error())
		}
	}))
	defer fixtureServer.Close()

	backendURL, err := url.Parse(fixtureServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	handler := &ProxyHandler{
		SiteAddress: backendURL,
	}

	srv := httptest.NewServer(handler)
	defer srv.Close()

	_, err = url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Error(err)
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf(err.Error())
		}
		bodyChan <- respBody
	}()

	respHeaders := <-headersChan
	respBody := <-bodyChan

	log.Println("resp.Header:", respHeaders)

	actual := string(respBody)
	if actual != expected {
		t.Errorf("Actual: %s expected: %s", actual, expected)
	}
}
