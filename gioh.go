package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

const (
	host   = "localhost"
	scheme = "http"
)

func runServer(port string) {
	addr := host + ":" + port
	log.Printf("Running server at %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(addr, nil)
}

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

func runClient(port string, reqsCount int) {
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
