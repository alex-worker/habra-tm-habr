package response_handler

import "net/http"

type HttpResponseHanlderInterface interface {
	Handle(w http.ResponseWriter, resp *http.Response) error
}
