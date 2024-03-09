package main

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/server"
	"flag"
	"log"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

var testIpAndPort bool = false

func init() {
	flag.BoolVar(&testIpAndPort, "testIpAndPort", false, "使用 localhost:8080")
}
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	flag.Parse()
	ip := "localhost"
	port := "8080"
	if testIpAndPort {
		ip = "0.0.0.0"
		port = "80"
	}
	server.APP.Start(ip, port)
}
