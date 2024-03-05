package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
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
