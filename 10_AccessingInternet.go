package main

import ("fmt"
		"net/http"
		"io/ioutil")

func main() {
	// unused variables are _ as in go unused variables will throw erroe
	resp, _ := http.Get("https://www.w3schools.com/xml/plant_catalog.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()

}

