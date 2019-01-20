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
