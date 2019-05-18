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

//since we had only one string to pull,we do it here only and reference loc tag og xml
// we also delete the overriding function  as that fetched us strings. Dude we already got strings here.

type SitemapIndexlean struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"urlset>url>news>news>title"`
	Keywords  []string `xml:"urlset>url>news>news>keywords"`
	Locations []string `xml:"urlset>url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndexlean
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
