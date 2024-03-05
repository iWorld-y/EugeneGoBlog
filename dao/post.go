package dao

import (
	"EugeneGoBlog/models"
	"html/template"
)

func GetPostByID(postID int) (models.Post, error) {
	row := DB.QueryRow("select * from goblog.blog_post where pid=?", postID)

	var post models.Post
	if err := row.Scan(
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
		return post, err
	}
	return post, nil
}

func GetPostMores(posts []models.Post) []models.PostMore {
	//posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(post.Content),
			CategoryId:   post.CategoryId,
			CategoryName: GetCategoryNameByID(post.CategoryId),
			UserId:       post.UserId,
			UserName:     GetUserNameByID(post.UserId),
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	return postMores
}
func GetPostPageByCategortID(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from goblog.blog_post where category_id=? limit ?,?", cid, page, pageSize)
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
func CountGetPostsByCategoryID(cid int) int {
	var count int
	_ = DB.QueryRow("select count(1) from goblog.blog_post where category_id=?", cid).Scan(&count)
	return count
}
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
