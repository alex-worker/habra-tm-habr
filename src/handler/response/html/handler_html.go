package html

import (
	"habra-tm-habr/src/handler/response/html/parser"
	"habra-tm-habr/src/handler/response/html/replacer"
	"habra-tm-habr/src/handler/utils/headers"
	"net/http"
)

func Handle(w http.ResponseWriter, resp *http.Response) error {
	myHtml, err := parser.BytesToHTML(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	parser.Update(myHtml, replacer.DoSomeTM)
	myBytes, err := parser.HTMLToBytes(myHtml)
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
