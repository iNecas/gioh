package gioh

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
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

	var (
		respChan = make(chan response)
		wg       = sync.WaitGroup{}
	)
	wg.Add(reqsCount + 1)

	sendRequest := func(reqNumber int) {
		defer wg.Done()

		r := response{reqNumber: reqNumber}

		httpResp, err := http.Get(target)
		if err != nil {
			r.result = "FAIL"
		} else {
			r.result = httpResp.Status
		}

		respChan <- r
	}

	for i := 0; i < reqsCount; i++ {
		go sendRequest(i)
	}

	go func () {
		for r := range respChan {
			responses = append(responses, r)
			if len(responses) == reqsCount {
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()

	formatResponses(responses)
}
