package main

import (
	"net/http"

	bd "github.com/SKatiyar/blevedemo"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../../public")))
	mux.HandleFunc("/search", bd.Search)

	http.ListenAndServe(":9090", mux)
}
