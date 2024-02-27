package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"EugeneGoBlog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{{
		Cid:  1,
		Name: "go",
	}}
	var posts = []models.PostMore{{
		Pid:          1,
		Title:        "go博客",
		Content:      "内容",
		UserName:     "Eugene",
		ViewCount:    123,
		CreateAt:     "2024-02-27",
		CategoryId:   1,
		CategoryName: "go",
		Type:         0,
	}}
	hr := &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	index := common.Template.Index
	index.WriteData(w, hr)
}
