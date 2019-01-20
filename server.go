package gioh

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunServer(port string) {
	addr := host + ":" + port
	log.Printf("Running server at %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		time.Sleep(time.Second)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(addr, nil)
}
