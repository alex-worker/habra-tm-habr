package main

import (
	"habra-tm-habr/src/handler"
	"habra-tm-habr/src/metrics"
	"log"
	"net/http"
	"net/url"
)

const profileAddress = ":9090"

const proxyAddress = ":8080"
const siteAddress = "http://habrahabr.ru"

func main() {

	go metrics.RunMetrics(profileAddress)

	log.Printf("Proxy address %v -> %v\n", proxyAddress, siteAddress)

	proxyUrl, err := url.Parse(siteAddress)
	if err != nil {
		panic(err)
	}

	myRequestProcessor := &handler.RequestProcessor{
		SiteAddress: proxyUrl,
	}

	myHandler := &handler.ProxyHandler{
		ProcessRequest: myRequestProcessor.Request,
	}

	err = http.ListenAndServe(proxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
