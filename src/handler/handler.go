package handler

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ProxyHandler struct {
	SiteAddress *url.URL
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Fatalf("Method not supported %s", r.Method)
	}

	log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

	cli := http.Client{}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = p.SiteAddress.Host
	r.URL.Host = p.SiteAddress.Host
	r.URL.Scheme = p.SiteAddress.Scheme
	r.RequestURI = ""

	delHeaders(r.Header)
	resp, err := cli.Do(r)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(resp.Body)

	var respHandler func(w http.ResponseWriter, resp *http.Response) error

	contentType, _ := getContentType(resp.Header)
	if strings.HasPrefix(contentType, "text/html") {
		respHandler = handleHTML
	} else {
		respHandler = handleRaw
	}

	err = respHandler(w, resp)
	if err != nil {
		log.Println(err.Error())
	}
}
