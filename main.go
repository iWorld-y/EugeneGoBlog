package main

import (
	"EugeneGoBlog/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	path, _ := os.Getwd()
	fmt.Printf("Path:\t%s\n", path)
	home := path + "/template/home.html"
	footer := path + "/template/layout/footer.html"
	header := path + "/template/layout/header.html"
	personal := path + "/template/layout/personal.html"
	post_list := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post_list, pagination)

	hr := &models.HomeResponse{}
	if err := t.Execute(w, hr); err != nil {
		log.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
