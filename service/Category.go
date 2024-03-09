package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"log"
	"math"
)

func GetPostsByCategoryId(cid, page, pageSize int) (*models.CategoryResponse, error) {
	/*
		获取指定类别的文章
		cid: 文章类别 id
		page: 页数
		pageSize: 一页中的文章数量
	*/
	//页面上涉及到的所有的数据，必须有定义
	allCategories, err := dao.GetAllCategory()
	if err != nil {
		log.Println("GetAllIndexInfo error:\t", err)
		return nil, err
	}

	posts, err := dao.GetPostPageByCategortID(cid, page, pageSize)
	for i, _ := range posts {
		posts[i].Content = string([]rune(posts[i].Content)[:100])
	}
	postMores := dao.Post2PostMores(posts)

	total := dao.CountGetPostsByCategoryID(cid)
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
	categoryName := dao.GetCategoryNameByID(cid)
	categoryResponse := &models.CategoryResponse{
		HomeResponse: hr,
		CategoryName: categoryName,
	}
	return categoryResponse, nil
}
