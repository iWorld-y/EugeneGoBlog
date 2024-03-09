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
	postMore := dao.Post2PostMores([]models.Post{post})[0]
	postResponse := &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postResponse, nil
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
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
func SearchPost(condition string) (searchResponse []models.SearchResponse) {
	posts := dao.GetPostSearch(condition)
	for _, post := range posts {
		searchResponse = append(searchResponse, models.SearchResponse{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return
}
