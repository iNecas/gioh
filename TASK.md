# Milestone 7: Throttled Http Proxy

**1.** In `proxy.go` define new type satisfying `net/http.RoundTripper` interface:

```go
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
```


For more details:
- [net/http.RoundTripper interface](https://golang.org/pkg/net/http/#RoundTripper) ([local](http://localhost:6060/pkg/net/http/#RoundTripper))

**2.** in `proxy.go`, update `RunProxy` function to use our `roundTripper` for
`Transport`

```go
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
```

For more details:
- [net/http/httputil.ReverseProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy) ([local](http://localhost:6060/pkg/net/http/httputil/#ReverseProxy))


**3.** Add new parameter `proxyConcurrency` into `cmd/gioh.go`

```go
func main() {
	var mode string
	serverPort := "3001"
	proxyPort := "3002"
	clientReqsCount := 10
	proxyConcurrency := 2

	flag.StringVar(&mode, "mode", "", "Mode to run the script in. Possible values: server, proxy, client")
	flag.StringVar(&serverPort, "server-port", serverPort, "A port to run the server at.")
	flag.StringVar(&proxyPort, "proxy-port", proxyPort, "A port to run the proxy at.")
	flag.IntVar(&proxyConcurrency, "proxy-concurrency", proxyConcurrency, "Number of concurrent requests from proxy to server.")
	flag.IntVar(&clientReqsCount, "client-requests-count", clientReqsCount, "Number of requests for client to send.")
	flag.Parse()

	switch mode {
	case "client":
		gioh.RunClient(proxyPort, clientReqsCount)
	case "server":
		gioh.RunServer(serverPort)
	case "proxy":
		gioh.RunProxy(proxyPort, serverPort, proxyConcurrency)
	default:
		fmt.Printf("Unknown mode %s\n", mode)
		flag.Usage()
		os.Exit(1)
	}
}
```

**4.** Add sleep into `HandleFunc` in `server.go` to make the throttling more
visible:

```go
import (
	"fmt"
	"log"
	"net/http"
	"time"
)
```

```go
http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	log.Printf("Handling request at %s", req.URL.Path)
	time.Sleep(time.Second)
	fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
})
```

For more details:
- [time package](https://golang.org/pkg/time) ([local](http://localhost:6060/pkg/time))

**3.** Re-run the proxy:

```
go run cmd/gioh.go --mode proxy
```

**4.** Try that the client still works:

```
go run cmd/gioh.go --mode client
```

**5.** Switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m8
cat TASK.md
```
