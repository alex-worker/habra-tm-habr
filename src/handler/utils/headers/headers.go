package headers

import (
	"errors"
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

func CopyHeaders(dst, src http.Header) {
	for key, value := range src {
		for _, value2 := range value {
			dst.Add(key, value2)
		}
	}
}
