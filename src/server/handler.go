package server

import (
	nodes "habra-tm-habr/src/nodes"
	"habra-tm-habr/src/replacer"
	"io"
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

	contentType, _ := GetContentType(resp.Header)

	// HasPrefix for "text/html; charset=utf-8" case
	if strings.HasPrefix(contentType, "text/html") {
		myHtml, err := nodes.BytesToHTML(resp.Body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		nodes.NodeAddTM(myHtml, replacer.DoSomeTM)
		myBytes, err := nodes.HTMLToBytes(myHtml)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		DelHeaders(resp.Header)
		copyHeaders(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		_, err = w.Write(myBytes)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		respRaw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		copyHeaders(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		_, err = w.Write(respRaw)
		if err != nil {
			//log.Println(err.Error())
			return
		}
	}
}
