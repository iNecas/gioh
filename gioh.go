package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const host = "localhost"

func runServer(port string) {
	addr := host + ":" + port
	log.Printf("Running server at %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(addr, nil)
}

func runClient(port string, reqsCount int) {
	fmt.Printf("I wish to run %d requests against port %s\n", reqsCount, port)
}

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
		runClient(serverPort, clientReqsCount)
	case "server":
		runServer(serverPort)
	default:
		fmt.Printf("Unknown mode %s\n", mode)
		flag.Usage()
		os.Exit(1)
	}
}
