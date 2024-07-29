package request_handler

import (
	"errors"
	"fmt"
	"habra-tm-habr/internal/pkg/html/headers"
	"net/http"
	"net/url"
)

type RequestProxyHandler struct {
	siteUrl *url.URL
}

func NewRequestProxyHandler(siteAddress string) RequestProxyHandler {
	siteUrl, err := url.Parse(siteAddress)
	if err != nil {
		panic(err)
	}
	return RequestProxyHandler{siteUrl}
}

func (h *RequestProxyHandler) Request(r *http.Request) (*http.Response, error) {

	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("Method not supported %s\n", r.Method)
		return nil, errors.New(msg)
	}

	cli := http.Client{}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = h.siteUrl.Host
	r.URL.Host = h.siteUrl.Host
	r.URL.Scheme = h.siteUrl.Scheme
	r.RequestURI = ""

	headers.DelHeaders(r.Header)
	return cli.Do(r)
}
