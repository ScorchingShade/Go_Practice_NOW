package main

//heads up, go does not have a while loop

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//What is unmarshalling and marshalling?
/*In computer science, unmarshalling or unmarshaling refers to the process of transforming a representation of an object that was used for storage or transmission to a representation of the object that is executable.
A serialized object which was used for communication can not be processed by a computer program.
An unmarshalling interface takes the serialized object and transforms it into an executable form. Unmarshalling (similar to deserialization) is the reverse process of marshalling.


In few words, "marshalling" refers to the process of converting the data or the objects inbto a byte-stream, and "unmarshalling" is the reverse process of converting the byte-stream beack to their original data or object.
The conversion is achieved through "serialization".
*/

//since the washingtonpost daily changes there xml structure, we kinda pivoted here, we are no longer getting data from internet, but doing the same thing, we are parsing the data from an xml sheet

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
	//we make a location type array to store data from sitemap>location
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	//We create a second datatype news to store titles, keywords and locations
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	//This is to iterate through our datatype values
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndexlean
	var n News
	news_map := make(map[string]NewsMap)

	bytes := washpostXML
	//converting a byte stream to actual xml, so as to get values from it. Also we need to store the data in byte stream because that's what we get from internet when we get any xml.
	xml.Unmarshal(bytes, &s)

	//iterating through every index in the range of available urls/locations, we get locations from our sitemap
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		//for every location that we get, we need to be able to access it, go inside it
		for idx, _ := range n.Titles {

			//we use the map with the key as the titles, in the titles array, we move through locations like Title[0], Title[1] etc. We do this using the idx that we created.
			//idx itself iterates from 0 to the length of the range of titles

			//now this last part is a bit confusing
			//The Newsmap here is a structure, remember when we made a map on line 99? we defined the key as string, which here is n.Titles[idx], an array of strings
			//and the value as Newsmap, which is a stuct and can take many values.
			//so for every key, we are storing two values.
			// the map will look like ["TechnologyTitle1: {keywords,location of first news}, TechnologyTitle2:{ keywords,location of second news } ] and so on
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}

		}
	}

	//We simply iterate through our populated news_map  by proper formatting here
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
