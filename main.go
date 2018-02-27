package main

import (
	"net/http/httputil"
	"net/url"
	"log"
	"net/http"
	"os"
	"fmt"
)

func main() {

	upstreamURL := os.Getenv("UPSTREAM")

	fmt.Printf("Starting proxy for %s\n", upstreamURL)

	upstream, err := url.Parse(upstreamURL)

	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(upstream)

	proxy.Director = func(r *http.Request) {
		r.Host = upstream.Host
		r.URL.Host = r.Host
		r.URL.Scheme = upstream.Scheme
	}

	http.ListenAndServe(":8080", http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {

		proxy.ServeHTTP(wr, r)

	}))

}