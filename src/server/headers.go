package server

import (
	"errors"
	"log"
	"net/http"
)

var hopHeaders = []string{
	"Connection",
	"Content-Length",
	"Accept-Encoding",

	"Content-Encoding",
	"Transfer-Encoding",
	"Public-Key-Pins",
	"Keep-Alive",

	//"Upgrade-Insecure-Requests",
	//"Upgrade",
	//"Proxy-Authenticate",
	//"Proxy-Authorization",
	//"Te", // canonicalized version of "TE"
	//"Trailers",
	//"Transfer-Encoding",
	//"Sec-Fetch-Dest",
	//"Sec-Fetch-Mode",
	//"Sec-Fetch-Site",
	//"Accept-Encoding",
	//"Content-Encoding",
}

func DelHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func GetContentType(header http.Header) (string, error) {
	for key, value := range header {
		if key == "Content-Type" {
			return value[0], nil
		}
	}
	return "", errors.New("Content-Type not found")
}

func ShowHeaders(header http.Header) {
	for key, value := range header {
		//for _, value2 := range value {
		log.Println(key, ": ", value)
		//}
	}
}

func copyHeaders(dst, src http.Header) {
	for key, value := range src {
		for _, value2 := range value {
			dst.Add(key, value2)
		}
	}
}
