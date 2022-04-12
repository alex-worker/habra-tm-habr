package server

import (
	nodes "habra-tm-habr/src/nodes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ProxyHandler struct {
	siteAddress *url.URL
}

func newProxyHandler(proxyUrl *url.URL) http.Handler {
	handler := ProxyHandler{
		siteAddress: proxyUrl,
	}
	return handler
}

func (p ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Fatalf("Method not supported %s", r.Method)
	}

	log.Println("Addr: ", r.RemoteAddr, "Method:", r.Method, "URL: ", r.URL.String())

	cli := http.Client{}

	originServerURL := p.siteAddress
	//originServerURL, err := url.Parse("https://habrahabr.ru/")

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = originServerURL.Host
	r.URL.Host = originServerURL.Host
	r.URL.Scheme = originServerURL.Scheme
	r.RequestURI = ""

	DelHeaders(r.Header)
	resp, err := cli.Do(r)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	contentType, _ := GetContentType(resp.Header)
	respRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// HasPrefix for "text/html; charset=utf-8" case
	if strings.HasPrefix(contentType, "text/html") {
		strBody, _ := nodes.AddSomeTM(respRaw)
		copyHeaders(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		_, err := w.Write([]byte(strBody))
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		copyHeaders(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		_, err := w.Write(respRaw)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
