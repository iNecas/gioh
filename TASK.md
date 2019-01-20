# Milestone 3: Basic HTTP Client

**1.** Add `"text/tabwriter"` to the imports:

```go
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)
```

**2.** Add new constant `scheme`

```go
const (
	host   = "localhost"
	scheme = "http"
)
```

**3.** Define a new struct type to collect the client results

```go
type response struct {
	reqNumber int
	result string
}
```

For more details:
- [struct](https://tour.golang.org/moretypes/2)



**4.** define a function `formatResponses` that will take a slice of responses and
   print them to the stdout as a table using `tabwriter` package.

```go
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
```

For more details:

- [slice](https://golang.org/doc/effective_go.html#slices) ([local](http://localhost:6060/doc/effective_go.html#slices))
- [text/tabwriter package](https://godoc.org/text/tabwriter) ([local](http://localhost:6060/pkg/text/tabwriter/))


**5.** implement the `runClient` function that will make several requests to the server:

```go
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
```

For more details:

- [append function](https://golang.org/pkg/builtin/#append) ([local](http://localhost:6060/pkg/builtin/#append/))
- [net/http package](https://godoc.org/net/http) ([local](http://localhost:6060/pkg/net/http/))


**6.** Switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m4
cat TASK.md
```
