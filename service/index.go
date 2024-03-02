package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"log"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	//页面上涉及到的所有的数据，必须有定义
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("GetAllIndexInfo error:\t", err)
		return nil, err
	}
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
	return hr, nil
}
