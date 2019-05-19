package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var washpostXMLCom = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-technology-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
xmlns:n="http://www.google.com/schemas/sitemap-news/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd
http://www.google.com/schemas/sitemap-news/0.9
http://www.google.com/schemas/sitemap-news/0.9/sitemap-news.xsd">
<url>
<loc>https://www.washingtonpost.com/business/technology/un-adds-32-items-to-list-of-prohibited-goods-for-north-korea/2017/10/23/5f112818-b812-11e7-9b93-b97043e57a22_story.html</loc>
<changefreq>hourly</changefreq>
<n:news>
<n:publication>
<n:name>Washington Post</n:name>
<n:language>en</n:language>
</n:publication>
<n:publication_date>2017-10-23T22:12:20Z</n:publication_date>
<n:title>UN adds 32 items to list of prohibited goods for North Korea</n:title>
<n:keywords>
UN-United Nations-North Korea-Sanctions,North Korea,East Asia,Asia,United Nations Security Council,United Nations,Business,General news,Sanctions and embargoes,Foreign policy,International relations,Government and politics,Government policy,Military technology,Technology</n:keywords>
</n:news>
</url>
<url>
<loc>https://www.washingtonpost.com/business/technology/cisco-systems-buying-broadsoft-for-19-billion-cash/2017/10/23/ae024774-b7f2-11e7-9b93-b97043e57a22_story.html</loc>
<changefreq>hourly</changefreq>
<n:news>
<n:publication>
<n:name>Washington Post</n:name>
<n:language>en</n:language>
</n:publication>
<n:publication_date>2017-10-23T21:42:14Z</n:publication_date>
<n:title>Cisco Systems buying BroadSoft for $1.9 billion cash</n:title>
<n:keywords>
US-Cisco-BroadSoft-Acquisition,Cisco Systems Inc,Business,Technology,Communication technology</n:keywords>
</n:news>
</url>
</urlset>`)

type SitemapIndexCom struct {
	//we make a location type array to store data from sitemap>location
	Locations []string `xml:"sitemap>loc"`
}

type NewsCom struct {
	//We create a second datatype news to store titles, keywords and locations
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMapCom struct {
	//This is to iterate through our datatype values
	Keyword  string
	Location string
}

//generating a struct to pass data to the html files
type NewsAggPageCom struct {
	Title string
	News  map[string]NewsMapCom
}

//a basic index page , this is not a template
func index_handlerCom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 align=\"center\">Templating</h1>")
}

func newsAggHandlerCom(w http.ResponseWriter, r *http.Request) {

	var s SitemapIndexCom
	var n NewsCom
	news_map := make(map[string]NewsMapCom)
	bytes := washpostXMLCom
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMapCom{n.Keywords[idx], n.Locations[idx]}

		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("/root/Desktop/Projects/Go_Pros/Practice/Prac15_Basics_basictemplating.html")

	fmt.Println(t.Execute(w, p))
}

func main() {

	http.HandleFunc("/", index_handlerCom)
	http.HandleFunc("/agg", newsAggHandlerCom)
	http.ListenAndServe(":8000", nil)

}
