// Simple Reverse proxy
// 1/2/2017 - Nick Vellios

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		director := func(req *http.Request) {
			req = r
			req.URL.Scheme = "http"
			req.URL.Host = "somehost.com" //r.Host
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServeTLS(":443", "/path/to/apache.crt", "/path/to/apache.key", nil))
}
