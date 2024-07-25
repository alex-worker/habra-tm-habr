package request_handler

import "net/http"

type HttpRequestHadnlerInterface interface {
	Request(r *http.Request) (*http.Response, error)
}
