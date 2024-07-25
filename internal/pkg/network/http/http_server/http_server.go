package http_server

import (
	"log"
	"net/http"
)

type HttpServerInterface interface {
	Run(listenAddress string, h http.Handler) error
}

type HttpServer struct{}

func NewHttpServer() HttpServerInterface {
	return &HttpServer{}
}

func (s *HttpServer) Run(listenAddress string, h http.Handler) error {
	log.Printf("HttpServer run...")

	err := http.ListenAndServe(listenAddress, h)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}
