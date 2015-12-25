package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	// Posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", PostsShowHandler)
	r.PUT("/posts/:id", PostsUpdateHandler)
	r.GET("/posts/:id/edit", PostsEditHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":"+port, r)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

func PostsShowHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

func PostsUpdateHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "update post", id)
}

func PostsEditHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
