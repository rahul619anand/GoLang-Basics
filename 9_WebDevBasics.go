package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats based on type
	fmt.Fprintf(w, "<h1>Go</h1>")
	fmt.Fprintf(w, "<p>Goa</p>")
	fmt.Fprintf(w, "<p>Gone</p>")
	fmt.Fprintf(w, "<p>Hello %s</p>","rahul")

	fmt.Fprintf (w,`<h1> Multi </h1>
		<p> line </p>
		<p> comment </p>`)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8080", nil)
}
