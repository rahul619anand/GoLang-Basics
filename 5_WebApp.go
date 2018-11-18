package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats based on type
	fmt.Fprintf(w, "Go Goa Gone")
}


func about_handler(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats based on type
	fmt.Fprintf(w, "Learning GoLang")
}


func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8080", nil)
}

