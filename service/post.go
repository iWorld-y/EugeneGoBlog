package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"log"
)

func GetPostDetail(postID int) (*models.PostRes, error) {
	post, err := dao.GetPostByID(postID)

	if err != nil {
		return nil, err
	}
	postMore := dao.GetPostMores([]models.Post{post})[0]
	postResponse := &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postResponse, nil
}

func Writing() (writingResponse models.WritingResponse) {
	writingResponse.Title = config.Cfg.Viewer.Title
	writingResponse.CdnURL = config.Cfg.System.CdnURL

	categoies, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	writingResponse.Categories = categoies
	return
}
