package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/nathandennis/go-proxy/functions"
)

func main() {
	handler := &httputil.ReverseProxy{}
	handler.Director = functions.HandlerDirector

	s := &http.Server{
		Addr:         ":1412",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Proxy server listening on %s\n", s.Addr)
	s.ListenAndServe()
}
