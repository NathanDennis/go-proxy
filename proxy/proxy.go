package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// NewReverseProxyHandler sets the handler.Director to the proxyDirector func
func NewReverseProxyHandler() http.Handler {
	return &httputil.ReverseProxy{
		Director: proxyDirector,
	}
}

// SingleJoiningSlash determines whether a slash should be added
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

// JoinURLPath is the same as SingleJoiningSlash, but uses EscapedPath to determine
// whether a slash should be added
func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}

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

// proxyDirector contains the logic for handler.Director in main() of main.go
func proxyDirector(req *http.Request) {
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
