package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	indexData := IndexData{
		Title: "Eugene 的博客",
		Desc:  "入门学习作业",
	}
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHTML(w http.ResponseWriter, r *http.Request) {
	indexData := IndexData{
		Title: "Eugene 的博客",
		Desc:  "入门学习作业",
	}
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHTML)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
