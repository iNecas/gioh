# Milestone 5: Concurrent Requests

**1.** Add `sync` package into imports of `client.go` file:

```go
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"sync"
)
```

**2.** In `client.go` file, update the `RunClient` function to use channels to
send the requests concurrently:

```go
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
```

For more details:
- [Sync.WaitGroup struct](https://golang.org/pkg/sync/#WaitGroup) ([local](http://localhost:6060/pkg/sync/#WaitGroup))
- [channels](https://golang.org/doc/effective_go.html#channels) ([local](http://localhost:6060/doc/effective_go.html#channels))
- [make function](https://golang.org/pkg/builtin/#make) ([local](http://localhost:6060/pkg/builtin/#make))
- [goroutines](https://golang.org/doc/effective_go.html#goroutines) ([local](http://localhost:6060/doc/effective_go.html#goroutines))
- [defer](https://golang.org/doc/effective_go.html#defer) ([local](http://localhost:6060/doc/effective_go.html#defer))


**3.** Try that the client still works:

```
go run cmd/gioh.go --mode client
```

**4.** Switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m6
cat TASK.md
```
