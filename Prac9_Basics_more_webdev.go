package main

import (
	"fmt"
	"net/http"
)

func index1_handler(w http.ResponseWriter, r *http.Request) {

	//we can write anything like html in this
	fmt.Fprintf(w, "<h1 align=\"center\" style=\"color:red\">This is web</h1>")

	fmt.Fprintf(w, "<p align=\"center\">Making different text appear</p>"+
		"<p align=\"center\">Making different text appear</p>"+
		"<p align=\"center\">Making different text appear</p>"+
		"moving up %s", "<strong>Best</strong>")
}

func main() {
	http.HandleFunc("/", index1_handler)
	http.ListenAndServe(":8000", nil)
}
