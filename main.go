package main

import (
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func main() {
	restConf := &rest.RestConf{}
	conf.MustLoad("etc/helloworld.yaml", restConf)
	s, err := rest.NewServer(*restConf)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			httpx.OkJson(rw, "Hello, World! ")
		},
	})
	defer s.Stop()
	s.Start()
}
