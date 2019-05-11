package main

//we need main package and fmt or format import to print anything basically.
//GO isn't object oriented , does not have classes and shit, so like make functions.

import (
	"fmt"
	"net/http"
)

//The index_handler function contains the data of web page to load.
// http.ResponseWriter is a custom data type, *http.Request is not pointer but actually reading through the request
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This is a webpage</h1>")
}

//This fuction initialise our webapp
//The HandleFunc of http requires net/http library
//Handlefunc gives the starting index of webpage, i.e the address.
//the ListenAndServe opens the port to present our webpage.
func webApp() {
	http.HandleFunc("/admin", index_handler)
	http.ListenAndServe(":8000", nil)
}

func main() {
	webApp()
}
