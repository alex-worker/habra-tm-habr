package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type IRequestProcessor interface {
	Request(r *http.Request) (*http.Response, error)
}

type RequestProcessor struct {
	SiteAddress *url.URL
}

func (h *RequestProcessor) Request(r *http.Request) (*http.Response, error) {

	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("Method not supported %s\n", r.Method)
		return nil, errors.New(msg)
	}

	cli := http.Client{}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = h.SiteAddress.Host
	r.URL.Host = h.SiteAddress.Host
	r.URL.Scheme = h.SiteAddress.Scheme
	r.RequestURI = ""

	delHeaders(r.Header)
	return cli.Do(r)
}
