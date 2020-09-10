package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nathandennis/go-proxy/proxy"
)

func main() {
	s := &http.Server{
		Addr:         ":1412",
		Handler:      proxy.NewReverseProxyHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Proxy server listening on %s\n", s.Addr)
	s.ListenAndServe()
}
