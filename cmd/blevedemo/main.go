package main

import (
	"fmt"
	"net/http"

	bd "github.com/SKatiyar/blevedemo"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../../public")))
	mux.HandleFunc("/search", bd.Search)

	go func() {
		if indexErr := bd.Index("../../public/sites.csv"); indexErr != nil {
			panic(indexErr)
		}
		fmt.Println("Indexing completed")
	}()

	fmt.Println("Listening on :9090")
	http.ListenAndServe(":9090", mux)
}
