package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func main() {

	handler := &httputil.ReverseProxy{}
	handler.Director = func(req *http.Request) {
		dest, err := url.Parse(req.Header.Get("X-Proxy-Target"))
		if err != nil {
			fmt.Printf("Error parsing proxy target: %v\n", err)
		}

		req.URL.Path, req.URL.RawPath = joinURLPath(dest, req.URL)

		if dest.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = dest.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = dest.RawQuery + "&" + req.URL.RawQuery
		}

		req.Header.Del("X-Proxy-Target")
		req.URL.Host = dest.Host
		req.URL.Scheme = dest.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = dest.Host

		fmt.Println(req.URL)
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

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}
