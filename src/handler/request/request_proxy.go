package request

import (
	"errors"
	"fmt"
	"habra-tm-habr/src/handler/utils/headers"
	"net/http"
	"net/url"
)

func NewRequestProxy(proxyUrl *url.URL) ProcessRequest {
	return func(r *http.Request) (*http.Response, error) {
		if r.Method != http.MethodGet {
			msg := fmt.Sprintf("Method not supported %s\n", r.Method)
			return nil, errors.New(msg)
		}
		cli := http.Client{}
		r.Host = proxyUrl.Host
		r.URL.Host = proxyUrl.Host
		r.URL.Scheme = proxyUrl.Scheme
		r.RequestURI = ""

		headers.DelHeaders(r.Header)
		return cli.Do(r)
	}
}
