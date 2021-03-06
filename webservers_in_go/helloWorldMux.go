package main

import (
	"io"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}