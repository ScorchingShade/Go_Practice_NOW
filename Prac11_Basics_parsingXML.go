package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
//The sitemap we are going to parse

<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/local.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/sports.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/national.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/world.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/business.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/technology.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/lifestyle.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/entertainment.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
</loc>
</sitemap>
</sitemapindex>

*/

//encoding/xml is to stucture data as xml

//defining structure

type SitemapIndex struct {
	//we created an array of Location type struct
	//L of location must be capital so that export to xml works, its syntax man
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	//loc must be lower case as tag is itself lowercase in sitemao
	Loc string `xml:"loc"`
}

//anything with square brackets and number in it is an array, anything without is a slice. arrays are fixed, slices aren't
//e.g arr [5] is array of int type
//e.g arr [] is slice of whatever type it is assigned to

//overriding string class with our own method to use here, getting everything only as strings not struct
func (l Location) String() string {
	//we use a value receiver to get a string value from xml location tag with this function

	//Sprintf formats strings for us
	return fmt.Sprint(l.Loc)

}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

}

//this is how to access internet
