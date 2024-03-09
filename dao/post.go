package dao

import (
	"EugeneGoBlog/models"
	"database/sql"
	"html/template"
	"log"
)

func GetPostSearch(condition string) (posts []models.Post) {
	rows, err := DB.Query("select * from goblog.blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		log.Println(err)
		return nil
	}
	posts, err = readPostsFromRows(rows)
	if err != nil {
		log.Println(err)
		return nil
	}
	return posts
}

func GetPostByID(postID int) (post models.Post, err error) {
	row, err := DB.Query("select * from goblog.blog_post where pid=?", postID)
	if err != nil {
		log.Println(err)
		return models.Post{}, err
	}

	posts, err := readPostsFromRows(row)
	if err != nil {
		log.Println(err)
		return models.Post{}, err
	}
	post = posts[0]
	return post, nil
}

func Post2PostMores(posts []models.Post) (postMores []models.PostMore) {
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

func readPostsFromRows(rows *sql.Rows) (posts []models.Post, err error) {
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
func GetPostPageByCategortID(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from goblog.blog_post where category_id=? limit ?,?", cid, page, pageSize)
	if err != nil {
		return nil, err
	}
	posts, _ := readPostsFromRows(rows)

	return posts, nil
}
func CountGetPostsByCategoryID(cid int) (cnt int) {
	_ = DB.QueryRow("select count(1) from goblog.blog_post where category_id=?", cid).Scan(&cnt)
	return cnt
}
func CountGetAllPost() (cnt int) {
	_ = DB.QueryRow("select count(1) from goblog.blog_post").Scan(&cnt)
	return cnt
}

func CountGetPostPageBySlug(slug string) (cnt int) {
	_ = DB.QueryRow("select count(1) from goblog.blog_post where slug=?", slug).Scan(&cnt)
	return cnt
}

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from goblog.blog_post order by update_at desc")
	if err != nil {
		return nil, err
	}
	posts, _ := readPostsFromRows(rows)

	return posts, nil
}

func GetPostPageBySlug(slug string, page, pageSize int) (posts []models.Post, err error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from goblog.blog_post where slug=? limit ?,?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	if posts, err = readPostsFromRows(rows); err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}

func GetPostPage(page, pageSize int) (posts []models.Post, err error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from goblog.blog_post order by update_at DESC limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}

	if posts, err = readPostsFromRows(rows); err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}

func SavePost(post *models.Post) {
	result, err := DB.Exec("insert into goblog.blog_post (title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at) values (?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)
	if err != nil {
		log.Println("保存文章失败")
	}
	postID, _ := result.LastInsertId()
	post.Pid = int(postID)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update goblog.blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println("更新文章失败:\t", err)
	}
}
