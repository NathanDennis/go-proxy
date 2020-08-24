package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {

	handler := &httputil.ReverseProxy{}
	handler.Director = func(req *http.Request) {
		dest, err := url.Parse(req.Header.Get("X-Proxy-Target"))
		if err != nil {
			fmt.Printf("Error parsing proxy target: %v\n", err)
		}

		req.Header.Del("X-Proxy-Target")
		req.URL.Host = dest.Host
		req.URL.Scheme = dest.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = dest.Host
	}

	s := &http.Server{
		Addr:         ":1412",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Proxy server listening on %s\n", s.Addr)
	s.ListenAndServe()
}
