package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go is super easy!")
}

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)

}
