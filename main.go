package main

import (
	"fmt"
	"log"
	"net/http"
)

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", UrlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}

func main() {
    ...
    http.HandleFunc("/color", ColorHandler)
    ...
}
