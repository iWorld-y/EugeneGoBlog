package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"html/template"
	"log"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	//页面上涉及到的所有的数据，必须有定义
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("GetAllIndexInfo error:\t", err)
		return nil, err
	}

	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML([]rune(post.Content)[:100]),
			CategoryId:   post.CategoryId,
			CategoryName: dao.GetCategoryNameByID(post.CategoryId),
			UserId:       post.UserId,
			UserName:     dao.GetUserNameByID(post.UserId),
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	hr := &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		1,
		1,
		[]int{1},
		true,
	}
	return hr, nil
}
