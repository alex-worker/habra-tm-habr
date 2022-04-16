package main

import (
	"habra-tm-habr/src/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
	"net/url"
)

const profileAddress = ":9090"

const proxyAddress = ":8080"
const siteAddress = "http://habrahabr.ru"

func main() {

	go func() {
		log.Printf("Profile address: %v/debug/pprof/\n", profileAddress)
		err := http.ListenAndServe(profileAddress, nil)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	log.Printf("Proxy address %v -> %v\n", proxyAddress, siteAddress)

	proxyUrl, err := url.Parse(siteAddress)
	if err != nil {
		panic(err)
	}

	myRequestProcessor := handler.RequestProcessor{
		SiteAddress: proxyUrl,
	}

	myHandler := &handler.ProxyHandler{
		Processor: &myRequestProcessor,
	}

	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(proxyAddress, myHandler)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
