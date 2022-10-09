package main

import (
    "net/http"
    "html"
    "log"
    "flag"
    "fmt"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
    flag.Parse()

    http.HandleFunc("/", HelloWeb)

    log.Print("Serving.  Go to `http://localhost:8080/foo/bar` in your web browser.")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
