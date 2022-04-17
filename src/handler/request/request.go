package request

import "net/http"

type ProcessRequest func(r *http.Request) (*http.Response, error)
