package test

import (
	"habra-tm-habr/src/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ManyRequest(t *testing.T) {
	t.Log("Test_Main")

	myRequestProcessor := &RequestProcessorEmpty{}

	srv := httptest.NewServer(&handler.ProxyHandler{
		ProcessRequest: myRequestProcessor.Request,
	})
	defer srv.Close()

	for i := 0; i < 10; i++ {
		_, err := http.Get(srv.URL)
		if err != nil {
			t.Error(err)
		}
	}
}
