# Milestone 6: Basic HTTP Proxy

**1.** Add `proxy.go` file:

```go
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
```

For more details:
- [net/url.URL struct](https://golang.org/pkg/net/url/#URL) ([local](http://localhost:6060/pkg/net/url/#URL))
- [net/http/httputil.NewSingleHostReverseProxy](https://golang.org/pkg/net/http/httputil/#NewSingleHostReverseProxy) ([local](http://localhost:6060/pkg/net/http/httputil/#NewSingleHostReverseProxy))


**2.** Add new mode into `cmd/gioh.go`

```go
func main() {
	var mode string
	serverPort := "3001"
	proxyPort := "3002"
	clientReqsCount := 10

	flag.StringVar(&mode, "mode", "", "Mode to run the script in. Possible values: server, proxy, client")
	flag.StringVar(&serverPort, "server-port", serverPort, "A port to run the server at.")
	flag.StringVar(&proxyPort, "proxy-port", proxyPort, "A port to run the proxy at.")
	flag.IntVar(&clientReqsCount, "client-requests-count", clientReqsCount, "Number of requests for client to send.")
	flag.Parse()

	switch mode {
	case "client":
		gioh.RunClient(proxyPort, clientReqsCount)
	case "server":
		gioh.RunServer(serverPort)
	case "proxy":
		gioh.RunProxy(proxyPort, serverPort)
	default:
		fmt.Printf("Unknown mode %s\n", mode)
		flag.Usage()
		os.Exit(1)
	}
}
```

For more details:
- [Sync.WaitGroup struct](https://golang.org/pkg/sync/#WaitGroup) ([local](http://localhost:6060/pkg/sync/#WaitGroup))
- [channels](https://golang.org/doc/effective_go.html#channels) ([local](http://localhost:6060/doc/effective_go.html#channels))
- [make function](https://golang.org/pkg/builtin/#make) ([local](http://localhost:6060/pkg/builtin/#make))
- [goroutines](https://golang.org/doc/effective_go.html#goroutines) ([local](http://localhost:6060/doc/effective_go.html#goroutines))
- [defer](https://golang.org/doc/effective_go.html#defer) ([local](http://localhost:6060/doc/effective_go.html#defer))


**3.** Run the proxy:

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
git checkout m7
cat TASK.md
```
