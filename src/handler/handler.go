package handler

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ProxyHandler struct {
	siteAddress *url.URL
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Fatalf("Method not supported %s", r.Method)
	}

	log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

	cli := http.Client{}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = p.siteAddress.Host
	r.URL.Host = p.siteAddress.Host
	r.URL.Scheme = p.siteAddress.Scheme
	r.RequestURI = ""

	DelHeaders(r.Header)
	resp, err := cli.Do(r)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(resp.Body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	var respHandler func(w http.ResponseWriter, resp *http.Response) error

	contentType, err := GetContentType(resp.Header)
	if err != nil {
		respHandler = handleRaw
	} else if strings.HasPrefix(contentType, "text/html") {
		respHandler = handleHTML
	} else {
		respHandler = handleRaw
	}

	err = respHandler(w, resp)
	if err != nil {
		log.Println(err.Error())
	}
}

func NewProxyHandler(proxyUrl *url.URL) (http.Handler, error) {
	return &ProxyHandler{
		siteAddress: proxyUrl,
	}, nil
}
