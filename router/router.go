package router

import (
	"EugeneGoBlog/api"
	"EugeneGoBlog/views"
	"net/http"
)

type IndexData struct {
	Title string
	Desc  string
}

func Router() {
	http.HandleFunc("/", views.HTML.Index)

	// c: means Category, View articles in a given category.
	http.HandleFunc("/c/", views.HTML.Category)
	// the page of log in
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
