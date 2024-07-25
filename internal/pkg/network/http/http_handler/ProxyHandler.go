package http_handler

import (
	"habra-tm-habr/internal/pkg/html/headers"
	"habra-tm-habr/internal/pkg/network/http/request_handler"
	"habra-tm-habr/internal/pkg/network/http/response_handler"
	"io"
	"log"
	"net/http"
	"strings"
)

type ProxyHandler struct {
	handlerRequest request_handler.HttpRequestHadnlerInterface
	handlerRaw     response_handler.HttpResponseHanlderInterface
	handlerHtml    response_handler.HttpResponseHanlderInterface
}

func NewProxyHandler(
	handlerRequest request_handler.HttpRequestHadnlerInterface,
	handlerRaw response_handler.HttpResponseHanlderInterface,
	handlerHtml response_handler.HttpResponseHanlderInterface) http.Handler {

	return &ProxyHandler{
		handlerRequest: handlerRequest,
		handlerRaw:     handlerRaw,
		handlerHtml:    handlerHtml,
	}
}

func bodyClose(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf(err.Error())
	}
}

func (p ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

	resp, err := p.handlerRequest.Request(r)
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer bodyClose(resp.Body)

	var respHandler func(w http.ResponseWriter, resp *http.Response) error

	contentType, err := headers.GetContentType(resp.Header)
	if err != nil {
		err = p.handlerRaw.Handle(w, resp)
	} else if strings.HasPrefix(contentType, "text/html") {
		err = p.handlerHtml.Handle(w, resp)
	} else {
		err = p.handlerRaw.Handle(w, resp)
	}

	err = respHandler(w, resp)
	if err != nil {
		log.Println(err.Error())
	}

}
