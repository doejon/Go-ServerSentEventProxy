package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//Std http handler interface implementation
type SSEProxy interface {
	http.Handler
}

type sseProxy struct {
	*httputil.ReverseProxy
}

//Create a new sse proxy
func NewSSEProxy(proxyUrl string) (SSEProxy, error) {
	s := &sseProxy{}
	u, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}
	s.ReverseProxy = httputil.NewSingleHostReverseProxy(u)
	s.ReverseProxy.FlushInterval = 100 * time.Millisecond

	//Prolong timeouts
	s.ReverseProxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   24 * time.Hour,
			KeepAlive: 24 * time.Hour,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	return s, nil
}
