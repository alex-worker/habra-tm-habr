package response_handler

import (
	"habra-tm-habr/internal/pkg/html/headers"
	"habra-tm-habr/internal/pkg/html/nodes"
	"net/http"
)

type HandlerHtml struct {
	processor TextProcessor
}

type TextProcessor interface {
	ProcessText(str string) string
}

func NewHandlerHtml(processor TextProcessor) HandlerHtml {
	return HandlerHtml{
		processor,
	}
}

func (h *HandlerHtml) Handle(w http.ResponseWriter, resp *http.Response) error {
	myHtml, err := nodes.BytesToHTML(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	nodes.Update(myHtml, h.processor.ProcessText)
	myBytes, err := nodes.HTMLToBytes(myHtml)
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
