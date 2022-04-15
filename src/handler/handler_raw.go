package handler

import (
	"io"
	"net/http"
)

func handleRaw(w http.ResponseWriter, resp *http.Response) error {
	myBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	delHeaders(resp.Header)
	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(myBytes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
