package main

import (
	"github.com/abalone0204/building_webapp/Godeps/_workspace/src/github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	rw.Write(markdown)
}
