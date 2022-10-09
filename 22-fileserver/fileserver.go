package main

import (
	"flag"
	"log"
	"net/http"
)

var PubDir = flag.String("pubdir", ".", "Directory to publish")

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*PubDir)))

	log.Print("Serving.  Go to `http://localhost:8080/` in your web browser.")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
