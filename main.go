package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})))
}
