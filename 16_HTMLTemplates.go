package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAppPage struct {
	Title string 
	News map[string]
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAppPage{Title: "Amazing News Agg", News: "some news"}
	t, _ := template.ParseFiles("basictemplate.html")
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
