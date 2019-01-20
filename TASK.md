# Milestone 4: Package Structure

**1.** Extract the main function into separate package:

Create a file `cmd/gioh.go` and copy the main function over (with necessary imports)

```go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mycoolnick/gioh"
)

func main() {
	var mode string
	serverPort := "3001"
	clientReqsCount := 10

	flag.StringVar(&mode, "mode", "", "Mode to run the script in. Possible values: server, proxy, client")
	flag.StringVar(&serverPort, "server-port", serverPort, "A port to run the server at.")
	flag.IntVar(&clientReqsCount, "client-requests-count", clientReqsCount, "Number of requests for client to send.")
	flag.Parse()

	switch mode {
	case "client":
		gioh.RunClient(serverPort, clientReqsCount)
	case "server":
		gioh.RunServer(serverPort)
	default:
		fmt.Printf("Unknown mode %s\n", mode)
		flag.Usage()
		os.Exit(1)
	}
}
```

Note we needed to import new package `github.com/mycoolnick/gioh`, from which we
called `RunClient` and `RunServer` functions. Also note the capital letters on
those functions (this makes the functions public).

We needed to clean the unused imports, as go would complain otherwise.

**2.** Extract the server part into `server.go` file:

```go
package gioh

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port string) {
	addr := host + ":" + port
	log.Printf("Running server at %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(addr, nil)
}
```

**3.** Extract the client part into `client.go` file:

```go
package gioh

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

type response struct {
	reqNumber int
	result    string
}

func formatResponses(responses []response) {
	wr := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', 0)

	for _, s := range responses {
		fmt.Fprintf(wr, "%d\t%s\n", s.reqNumber, s.result)
	}

	err := wr.Flush()
	if err != nil {
		log.Fatalf("Error writing output: %s", err)
	}
}

func RunClient(port string, reqsCount int) {
	var responses []response
	target := scheme + "://" + host + ":" + port

	for i := 0; i < reqsCount; i++ {
		r := response{reqNumber: i}

		httpResp, err := http.Get(target)
		if err != nil {
			r.result = "FAIL"
		} else {
			r.result = httpResp.Status
		}
		responses = append(responses, r)
	}

	formatResponses(responses)
}
```

**4.** Extract the constants part into `const.go` file:

```go
package gioh

const (
	host   = "localhost"
	scheme = "http"
)
```

**5.** Delete the original `gioh.go` file

**6.** Try that the code still works:

```
go run cmd/gioh.go --mode server
go run cmd/gioh.go --mode client
```

**7.** Switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m5
cat TASK.md
```
