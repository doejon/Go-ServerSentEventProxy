package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	root, _ := url.Parse("http://localhost:3000/")
	sse, err := NewSSEProxy("http://localhost:3000/events/")
	if err != nil {
		panic(err)
	}

	http.Handle("/", httputil.NewSingleHostReverseProxy(root))

	http.Handle("/events/", sse)

	http.ListenAndServe(":10000", nil)
}
