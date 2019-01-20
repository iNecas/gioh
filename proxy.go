package gioh

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func RunProxy(proxyPort, targetPort string) {
	target := url.URL{
		Host:   host + ":" + targetPort,
		Scheme: scheme,
	}
	http.Handle("/", httputil.NewSingleHostReverseProxy(&target))
	log.Printf("Proxy started")
	http.ListenAndServe(":"+proxyPort, nil)
}
