# Go in one hour

The purpose of this repository (and related workshop) is using a real-world
example to guide you though basic principles of Go the programming language.


## Organization

The repository is split into multiple branches, each representing a specific
milestone. At every milestone, one is given a specific task to finish.

To start:

    git checkout m1
    cat TASK.md


## Before you start

Before you can start, you will need to [install Go environment](https://golang.org/doc/install). Also, it's
recommended to [enable your editor of choice](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) to edit Go files.


## Documentation

Since there might be occasions, where you don't have network connectivity, it's
good to have an option to read the documentation locally.

You can do so by running:

    godoc --http:6060 --play

Further in the milesones, we link to the documentation for more details about
the topics.

See [Go lang learning page](https://golang.org/doc/#learning) ([local](http://localhost:6060/doc/#learning)) for more generic introduction to Go syntax and
additional resources.


## Milestones

-   **[m1 - Intro](https://github.com/iNecas/gioh/blob/m1/TASK.md)**

    Basic go tools for developing. Initializing a go module. `fmt`, `log` and
    `flag` packages. Defining variables and functions.

-   **[m2 - Basic HTTP Server](https://github.com/iNecas/gioh/blob/m2/TASK.md)**

    Writing a simple HTTP server using `net/http` package. Using `log` and more `fmt`
    functions. Defining constants.

-   **[m3 - Basic HTTP Client](https://github.com/iNecas/gioh/blob/m3/TASK.md)**

    Writing an HTTP client using `net/http` package.

-   **[m4 - Package Structure](https://github.com/iNecas/gioh/blob/m4/TASK.md)**

    Splitting the code into various files across the `gioh` package.

-   **[m5 - Concurrent Requests](https://github.com/iNecas/gioh/blob/m4/TASK.md)**

    Expanding the HTTP client to perform multiple requests at a time. Basic
    usage of channels.

-   **[m6 - Basic HTTP Proxy](https://github.com/iNecas/gioh/blob/m5/TASK.md)**

    Writing a proxy to sit in frond of the HTTP server using `net.http.httputil` package.

-   **[m7 - Throttled Http Proxy](https://github.com/iNecas/gioh/blob/m6/TASK.md)**

    Limit the amount of requests to come from the proxy the to HTTP server by using channels.

-   **[m8 - Finished](https://github.com/iNecas/gioh/blob/m8/TASK.md)**

    The finalized example.
