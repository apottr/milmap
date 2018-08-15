package main

import (
	"log"
	"net/http"
)

func renderJS(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /bundle", r.Method)
	http.ServeFile(w, r, "js/dist/bundle")
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /", r.Method)
	http.ServeFile(w, r, "index.html")
}

func main() {
	port := ":8080"
	http.HandleFunc("/bundle", renderJS)
	http.HandleFunc("/", renderPage)
	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
