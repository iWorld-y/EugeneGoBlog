package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// printFileTree 打印文件树
func printFileTree(path string, depth int) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("无法读取目录 %s: %v\n", path, err)
		return
	}

	for _, file := range files {
		// 打印缩进和文件/目录名称
		fmt.Printf("%s%s\n", strings.Repeat("│   ", depth), file.Name())

		// 如果是目录，递归打印其内容
		if file.IsDir() {
			subDir := filepath.Join(path, file.Name())
			printFileTree(subDir, depth+1)
		}
	}
}
func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("无法获取当前目录: %v\n", err)
		return
	}

	fmt.Println(currentDir) // 打印当前目录路径
	printFileTree(currentDir, 0)
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
