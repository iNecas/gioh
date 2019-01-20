# Milestone 2: Basic HTTP Server

**1.** import new packages `"net/http"` and `"log"`

```go
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)
```

**2.** define a constant `host`, that will represent where our server is running

```go
const host = "localhost"
```

**3.** implement the `runServer` function:

```go
func runServer(port string) {
	addr := host + ":" + port
	log.Printf("Running server at %s", addr)
	
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(addr, nil)
}
```

For more details:

- [net/http package](https://godoc.org/net/http) ([local](http://localhost:6060/pkg/net/http/))
- [log package](https://godoc.org/log) ([local](http://localhost:6060/pkg/log/))


**4.** run the server `go run gioh.go --mode server`

**5.** in separate terminal, send a request against the server:

```
curl http://localhost:3001/hello/world
```

**6.** Switch to next milestone

```
git add -A .; git commit -m "My changes"
git checkout m3
cat TASK.md
```
