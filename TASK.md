# Milestone 1: Intro

**1.** initialize a new go module

```
go mod init github.com/mycoolnick/gioh
```

**2.** create a file `gioh.go` to print hello world

```go
package main

import (
"os"
"flag"
"fmt"
)

func runServer(port string) {
fmt.Printf("I wish to run a server at port %s\n", port)
}

func runClient(port string, reqsCount int) {
	fmt.Printf("I wish to run %d requests against port %s\n", reqsCount, port)
}

func main() {
	var mode string
	serverPort := "3001"
	clientReqsCount := 10

	flag.StringVar( &mode, "mode", "", "Mode to run the script in. Possible values: server, proxy, client" )
	flag.StringVar( &serverPort, "server-port", "3001", "A port to run the server at." )
	flag.IntVar(&clientReqsCount, "client-requests-count", clientReqsCount, "Number of requests for client to send.")
	flag.Parse()

	fmt.Printf("mode %s\n", mode)
	switch mode {
	case "client":
		runClient(serverPort, clientReqsCount)
	case "server":
		runServer( serverPort )
	default:
		fmt.Printf("Unknown mode %s\n", mode)
		flag.Usage()
		os.Exit(1)
	}
}
```

For more details:

- [flag package](https://godoc.org/flag) ([local](http://localhost:6060/pkg/flag/))
- [fmt package](https://godoc.org/fmt) ([local](http://localhost:6060/pkg/fmt/))
- [os package](https://godoc.org/os) ([local](http://localhost:6060/pkg/os/))

**3.** run `go run gioh.go`

**4.** run `go fmt`, notice how the `gioh.go` file changed

**5.** run `go build gioh.go` to compile the executable

**6.** run `./gioh --mode server`

**7.** switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m2
cat TASK.md
```
