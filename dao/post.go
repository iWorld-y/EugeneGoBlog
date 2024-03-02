package dao

import (
	"EugeneGoBlog/models"
)

func CountGetAllPost() int {
	var count int
	_ = DB.QueryRow("select count(1) from goblog.blog_post").Scan(&count)
	return count
}
func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from goblog.blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}
