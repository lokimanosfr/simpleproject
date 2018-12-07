package main

import (
	"net/http"
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Processing %s ....", r.URL.Path)
// }

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}
