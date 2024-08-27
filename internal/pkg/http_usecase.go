package pkg

import (
	"habra-tm-habr/internal/pkg/network/http/http_handler"
	"habra-tm-habr/internal/pkg/network/http/http_server"
	"habra-tm-habr/internal/pkg/network/http/request_handler"
	"habra-tm-habr/internal/pkg/network/http/response_handler"
	"net/http"
)

func NewHttpProxyTextProcessorHandler(siteAddress string, p response_handler.TextProcessor) http.Handler {
	handlerRequest := request_handler.NewRequestProxyHandler(siteAddress)
	handlerRaw := response_handler.NewHandlerRaw()
	handlerHttp := response_handler.NewHandlerHtml(p)

	myHandler := http_handler.NewProxyHandler(&handlerRequest, &handlerRaw, &handlerHttp)
	return myHandler
}

func NewHttpServer() http_server.HttpServerInterface {
	return http_server.NewHttpServer()
}
