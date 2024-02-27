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
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
