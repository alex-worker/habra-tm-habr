package handler

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type ProxyHandler struct {
	ProcessRequest func(r *http.Request) (*http.Response, error)
}

func bodyClose(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf(err.Error())
	}
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

	resp, err := p.ProcessRequest(r)
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer bodyClose(resp.Body)

	var respHandler func(w http.ResponseWriter, resp *http.Response) error

	contentType, err := getContentType(resp.Header)
	if err != nil {
		respHandler = handleRaw
	} else if strings.HasPrefix(contentType, "text/html") {
		respHandler = handleHTML
	} else {
		respHandler = handleRaw
	}

	err = respHandler(w, resp)
	if err != nil {
		log.Println(err.Error())
	}
}
