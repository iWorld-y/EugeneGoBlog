package service

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
)

func FindPostPigeonhole() models.PigeonholeResponse {
	// 查询所有的文章， 按照月份分组
	// 查询所有分类

	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		month := post.UpdateAt.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	categories, _ := dao.GetAllCategory()

	return models.PigeonholeResponse{
		config.Cfg.Viewer,
		config.Cfg.System,
		categories,
		pigeonholeMap,
	}
}
