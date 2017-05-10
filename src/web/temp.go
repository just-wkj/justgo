package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName)

	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
