package main

import (
	"habra-tm-habr/src/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
	"net/url"
)

const proxyAddress = ":8080"
const siteAddress = "http://habrahabr.ru"

func main() {

	go func() {
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	log.Println("Hello world!", proxyAddress, " -> ", siteAddress)

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

	err = http.ListenAndServe(proxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
