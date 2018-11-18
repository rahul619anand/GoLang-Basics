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
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	var s SitemapIndex

	bytes := washingtonPostXML
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	
	var n News
	news_map := make(map[string]NewsMap)

	// reading individual urls from washingtonPostXML
	for _,Location := range s.Locations {
		fmt.Printf("\n%s",Location)
		resp, _ := http.Get(Location)
		bytes,_ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes,&n)

		for idx, _ := range n.Titles{
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)

	}

}