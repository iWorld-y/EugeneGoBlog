package dao

import (
	"EugeneGoBlog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from goblog.blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询失败:\t", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:\t", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
