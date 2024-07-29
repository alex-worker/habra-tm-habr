package http_handler

import (
	"habra-tm-habr/internal/pkg/html/headers"
	"io"
	"log"
	"net/http"
	"strings"
)

type ProxyHandler struct {
	handlerRequest RequestHadnler
	handlerRaw     ResponseHanlder
	handlerHtml    ResponseHanlder
}

type RequestHadnler interface {
	Request(r *http.Request) (*http.Response, error)
}

type ResponseHanlder interface {
	Handle(w http.ResponseWriter, resp *http.Response) error
}

func NewProxyHandler(
	handlerRequest RequestHadnler,
	handlerRaw ResponseHanlder,
	handlerHtml ResponseHanlder) http.Handler {

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
