package main

//heads up, go does not have a while loop

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//basic for loop
// for loop without any parameter runs like while (true) or while (1) , that is forever
//check git commits to see wassup

type SitemapIndex1 struct {
	Locations []Location1 `xml:"sitemap"`
}

type Location1 struct {
	Loc string `xml:"loc"`
}

func (l Location1) String() string {
	//we use a value receiver to get a string value from xml location tag with this function

	//Sprintf formats strings for us
	return fmt.Sprint(l.Loc)

}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex1
	xml.Unmarshal(bytes, &s)

	//using loops to show different urls

	//this is a for each loop. the blank param is the index, we don't want to show index thus its blank
	// the location here is a complete new variable and gets assignment to a range of different values from the sitemap struct datatype var locations
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}

}
