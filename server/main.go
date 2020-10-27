package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Encoding", "gzip")
		}
		http.ServeFile(w, r, "./static"+r.URL.Path)
	})
	log.Fatalln(http.ListenAndServe(":8080", m))
}
