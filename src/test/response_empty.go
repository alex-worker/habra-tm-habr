package test

import (
	"net/http"
)

func RequestEmpty(*http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.Header = make(http.Header)
	//resp.Header.Set("Content-Type", "text/plain")
	resp.StatusCode = 200
	//resp.Body = ioutil.NopCloser(strings.NewReader("Lol"))
	resp.Body = http.NoBody
	return resp, nil
}
