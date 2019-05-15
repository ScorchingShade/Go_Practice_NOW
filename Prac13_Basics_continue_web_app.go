package main

//heads up, go does not have a while loop

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//we cleaned up the previous code to make it more lean
//Goal- visiting sitemap links

type SitemapIndexlean struct {
	//since we had only one string to pull,we do it here only and reference loc tag og xml
	// we also delete the overriding function  as that fetched us strings. Dude we already got strings here.
	Locations []string `xml:"sitemap>loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex1
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}

}
