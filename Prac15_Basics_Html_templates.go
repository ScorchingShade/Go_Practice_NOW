package main

//voila we have reached the HTML templating

import (
	"fmt"
	"html/template"
	"net/http"
)

//generating a struct to pass data to the html files
type NewsAggPage struct {
	Title string
	News  string
}

//a basic index page , this is not a template
func index_handlertemp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 align=\"center\">Templating</h1>")
}

//This makes a template . We use the Prac15_Basics_basictemplating.html to pass this data
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	//we create a variable with NewsAggPage dataType
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	//we use the template function with Parsefiles to parse this data to the template
	t, err := template.ParseFiles("/root/Desktop/Projects/Go_Pros/Practice/Prac15_Basics_basictemplating.html")
	fmt.Println(err)
	//The execute finally executes this
	t.Execute(w, p)
}

func main() {

	http.HandleFunc("/", index_handlertemp)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}
