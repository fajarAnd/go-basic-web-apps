package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello world")
	})

	barHandler := func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"Hello Bar")
	}

	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":3000", nil)
}
