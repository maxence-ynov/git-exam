package main

import (
	"log"
	"net/http"
)

func main() {
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
