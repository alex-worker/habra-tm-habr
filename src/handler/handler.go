package handler

import (
	"habra-tm-habr/src/handler/request"
	"habra-tm-habr/src/handler/response"
	ResponseHTML "habra-tm-habr/src/handler/response/html"
	ResponseRaw "habra-tm-habr/src/handler/response/raw"
	"habra-tm-habr/src/handler/utils/headers"
	"io"
	"log"
	"net/http"
	"strings"
)

type RequestToResponse func(w http.ResponseWriter, r *http.Request)

func bodyClose(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf(err.Error())
	}
}

func NewRequestToResponse(DoRequest request.ProcessRequest) RequestToResponse {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

		resp, err := DoRequest(r)
		if err != nil {
			log.Printf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer bodyClose(resp.Body)

		var myProcessResponse response.ProcessResponse

		contentType, err := headers.GetContentType(resp.Header)
		if err != nil {
			myProcessResponse = ResponseRaw.Handle
		} else if strings.HasPrefix(contentType, "text/html") {
			myProcessResponse = ResponseHTML.Handle
		} else {
			myProcessResponse = ResponseRaw.Handle
		}

		err = myProcessResponse(w, resp)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
