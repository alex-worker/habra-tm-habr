package server

import (
	"log"
	"net/http"
	"net/url"
)

func Listen(proxyAddress string, siteAddress string) *http.Server {
	log.Println("ProxyServer.Listen: ", proxyAddress)
	siteUrl, err := url.Parse(siteAddress)
	if err != nil {
		log.Fatalf(err.Error())
	}

	server := &http.Server{
		Addr:    proxyAddress,
		Handler: newProxyHandler(siteUrl),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return server
}
