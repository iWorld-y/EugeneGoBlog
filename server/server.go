package server

import (
	"EugeneGoBlog/router"
	"log"
	"net/http"
)

var APP = &EugeneGoBlogServer{}

type EugeneGoBlogServer struct {
}

func (app *EugeneGoBlogServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
