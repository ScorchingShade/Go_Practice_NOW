package main

//heads up, go does not have a while loop

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var washpostXML = []byte(`
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

//we cleaned up the previous code to make it more lean
//Goal- visiting sitemap links

//since we had only one string to pull,we do it here only and reference loc tag og xml
// we also delete the overriding function  as that fetched us strings. Dude we already got strings here.

type SitemapIndexlean struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndexlean
	var n News
	news_map := make(map[string]NewsMap)

	bytes := washpostXML
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
