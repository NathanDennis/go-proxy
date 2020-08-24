package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	handler := &httputil.ReverseProxy{}
	handler.Director = func(req *http.Request) {
		url := url.Parse(req.Header.Get("X-Proxy-Target"))
		req.Header.Del("X-Proxy-Target")
		req.URL.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = url.Host
	}

	//url := req.Header.Get("X-Proxy-Target")

}
