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
func GetCategoryNameByID(cid int) string {
	if row := DB.QueryRow("select name from goblog.blog_category where cid=?", cid); row.Err() == nil {
		var categoryName string
		_ = row.Scan(&categoryName)
		return categoryName
	} else {
		log.Println("GetCategoryNameByID 类别读取错误:\t", row.Err())
	}
	return ""
}
