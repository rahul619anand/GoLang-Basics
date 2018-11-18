package main

import ("fmt"
		"encoding/xml"
		"io/ioutil"
		"net/http")

var washingtonPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)



type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	bytes := washingtonPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	
	var n News
	for _,Location := range s.Locations {
		fmt.Printf("\n%s",Location)
		resp, _ := http.Get(Location)
		bytes,_ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes,&n)
	}

}