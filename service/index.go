package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"log"
	"math"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	//页面上涉及到的所有的数据，必须有定义
	allCategories, err := dao.GetAllCategory()
	if err != nil {
		log.Println("GetAllIndexInfo error:\t", err)
		return nil, err
	}

	var posts []models.Post
	var total int // 当前选择类别下的文章总数
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetPostPageBySlug(slug)
	}

	for i, _ := range posts {
		// 若文章正文长度大于 100, 则只显示前 100 个 unicode 字符, 其余省略
		if temp := len([]rune(posts[i].Content)); temp > 100 {
			posts[i].Content = string([]rune(posts[i].Content)[:100])
		}
	}
	postMores := dao.Post2PostMores(posts)

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
