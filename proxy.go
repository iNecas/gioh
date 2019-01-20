package gioh

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type roundTripper struct {
	bucket chan interface{}
}

func newRoundTripper(limit int) *roundTripper {
	bucket := make(chan interface{}, limit)
	tripper := roundTripper{bucket}

	for i := 0; i < limit; i++ {
		bucket <- nil
	}

	return &tripper
}

func (t *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("Waiting to proxy request %s", req.URL.Path)

	<-t.bucket
	defer func() {
		t.bucket <- nil
	}()

	log.Printf("Handling request %s", req.URL.Path)
	return http.DefaultTransport.RoundTrip(req)
}

func RunProxy(proxyPort, targetPort string, concurrency int) {
	target := url.URL{
		Host:   host + ":" + targetPort,
		Scheme: scheme,
	}
	singleHostProxy := httputil.NewSingleHostReverseProxy(&target)
	throttledTransport := newRoundTripper(concurrency)

	throttledProxy := httputil.ReverseProxy{
		Director:  singleHostProxy.Director,
		Transport: throttledTransport,
	}

	http.Handle("/", &throttledProxy)
	log.Printf("Proxy started")
	http.ListenAndServe(":"+proxyPort, nil)
}
