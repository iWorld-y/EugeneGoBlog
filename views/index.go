package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		index.WriteError(w, errors.New("系统内部错误\n"))
		return
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10

	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")

	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("(*HTMLApi) Index 首页获取数据失败:\t", err)
		index.WriteError(w, errors.New("系统内部错误\n"))
	}
	if err := index.WriteData(w, hr); err != nil {
		log.Println(err)
	}
}
