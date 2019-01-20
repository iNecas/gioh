package main

import (
	"flag"
	"fmt"
	"os"
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
