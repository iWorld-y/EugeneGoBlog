package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"log"
	"math"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	//页面上涉及到的所有的数据，必须有定义
	allCategories, err := dao.GetAllCategory()
	if err != nil {
		log.Println("GetAllIndexInfo error:\t", err)
		return nil, err
	}

	posts, err := dao.GetPostPage(page, pageSize)
	for i, _ := range posts {
		posts[i].Content = string([]rune(posts[i].Content)[:100])
	}
	postMores := dao.GetPostMores(posts)

	total := dao.CountGetAllPost()
	pageCount := int(math.Ceil(float64(total) / 10.0))
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	hr := &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: allCategories,
		Posts:      postMores,
		Total:      total,
		Page:       page,
		Pages:      pages,
		PageEnd:    page != pageCount,
	}
	return hr, nil
}
