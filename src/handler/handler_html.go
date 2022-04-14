package handler

import (
	"habra-tm-habr/src/nodes"
	"habra-tm-habr/src/replacer"
	"net/http"
)

func handleHTML(w http.ResponseWriter, resp *http.Response) error {
	myHtml, err := nodes.BytesToHTML(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	nodes.Update(myHtml, replacer.DoSomeTM)
	myBytes, err := nodes.HTMLToBytes(myHtml)
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
