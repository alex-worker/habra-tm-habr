package main

import (
	"habra-tm-habr/src/handler"
	"log"
	"net/http"
	"net/url"
)

const proxyAddress = ":8080"
const siteAddress = "http://habrahabr.ru"

func main() {
	log.Println("Hello world!", proxyAddress, siteAddress)

	proxyUrl, err := url.Parse(siteAddress)
	if err != nil {
		panic(err)
	}

	myHandler := &handler.ProxyHandler{
		SiteAddress: proxyUrl,
	}
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    proxyAddress,
		Handler: myHandler,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
