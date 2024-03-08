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

	// c: means Category, View posts in a given category.
	http.HandleFunc("/c/", views.HTML.Category)

	// p: means Post, View the entire post
	http.HandleFunc("/p/", views.HTML.Detail)

	// the page of log in
	http.HandleFunc("/login", views.HTML.Login)

	// the page of write blog
	http.HandleFunc("/writing", views.HTML.Writing)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiNiuToken)
	http.HandleFunc("/api/v1/login", api.API.Login)
	// CDN
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
