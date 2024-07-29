package response_handler

import (
	"habra-tm-habr/internal/pkg/html/headers"
	"io"
	"net/http"
)

type HandlerRaw struct{}

func NewHandlerRaw() HandlerRaw {
	return HandlerRaw{}
}

func (h *HandlerRaw) Handle(w http.ResponseWriter, resp *http.Response) error {
	myBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	headers.DelHeaders(resp.Header)
	headers.CopyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(myBytes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
