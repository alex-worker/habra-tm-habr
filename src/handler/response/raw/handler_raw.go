package raw

import (
	"habra-tm-habr/src/handler/headers"
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, resp *http.Response) error {
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
