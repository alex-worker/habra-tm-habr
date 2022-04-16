package test

import (
	"habra-tm-habr/src/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EmptyProcessor struct{}

func (e *EmptyProcessor) Request(*http.Request) (*http.Response, error) {
	resp := &http.Response{}
	return resp, nil
}

func Test_ManyRequest(t *testing.T) {
	t.Log("Test_Main")

	srv := httptest.NewServer(&handler.ProxyHandler{
		Processor: &EmptyProcessor{},
	})
	defer srv.Close()

}
