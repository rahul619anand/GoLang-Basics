package main

import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml")


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

type NewsAppPage struct {
	Title string 
	News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
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

	p := NewsAppPage{Title: "Amazing News Agg", News: news_map}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	t.Execute(w,p)
}


func index_handler(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats based on type
	fmt.Fprintf(w, "<h1>Go</h1>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
