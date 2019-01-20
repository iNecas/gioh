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
