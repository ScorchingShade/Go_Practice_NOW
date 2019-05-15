package main

//heads up, go does not have a while loop

import (
	"encoding/xml"
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

type News struct {
	titles    []string `xml:"url>news>title"`
	keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndexlean
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

	}

}
