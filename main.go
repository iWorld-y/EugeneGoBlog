package main

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/router"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

var port80 bool = false

func init() {
	flag.BoolVar(&port80, "port80", false, "使用 80 端口")
}
func main() {
	flag.Parse()

	port := 8080
	if port80 {
		port = 80
	}

	server := http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", port),
		//Addr: fmt.Sprintf("0.0.0.0:%d", port),
	}
	router.Router()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
