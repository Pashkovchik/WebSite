package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
	
}
