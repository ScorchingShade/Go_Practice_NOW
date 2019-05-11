package main

//we need main package and fmt or format import to print anything basically.
//GO isn't object oriented , does not have classes and shit, so like make functions.

import (
	"fmt"
	"net/http"
)

// http.ResponseWriter is a custom data type, *http.Request is not pointer but actually reading through the request
func index_handler(w http.ResponseWriter, r *http.Request)

func webApp() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}

func main() {
	webApp()
}
