package request

import (
	"errors"
	"fmt"
	"habra-tm-habr/src/handler/headers"
	"net/http"
	"net/url"
)

type Processor struct {
	SiteAddress *url.URL
}

func (h *Processor) Request(r *http.Request) (*http.Response, error) {

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

	headers.DelHeaders(r.Header)
	return cli.Do(r)
}
