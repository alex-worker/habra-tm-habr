package response

import "net/http"

type ProcessResponse func(w http.ResponseWriter, resp *http.Response) error
