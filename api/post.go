package api

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"EugeneGoBlog/service"
	"EugeneGoBlog/utils"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`        // 文章ID
	Title      string    `json:"title"`      // 文章ID
	Slug       string    `json:"slug"`       // 自定也页面 path
	Content    string    `json:"content"`    // 文章的html
	Markdown   string    `json:"markdown"`   // 文章的Markdown
	CategoryId int       `json:"categoryId"` //分类id
	UserId     int       `json:"userId"`     //用户id
	ViewCount  int       `json:"viewCount"`  //查看次数
	Type       int       `json:"type"`       //文章类型 0 普通，1 自定义文章
	CreateAt   time.Time `json:"createAt"`   // 创建时间
	UpdateAt   time.Time `json:"updateAt"`   // 更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章ID
	Slug         string        `json:"slug"`         // 自定也页面 path
	Content      template.HTML `json:"content"`      // 文章的html
	CategoryId   int           `json:"categoryId"`   // 文章的Markdown
	CategoryName string        `json:"categoryName"` // 分类名
	UserId       int           `json:"userId"`       // 用户id
	UserName     string        `json:"userName"`     // 用户名
	ViewCount    int           `json:"viewCount"`    // 查看次数
	Type         int           `json:"type"`         // 文章类型 0 普通，1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchResponse struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}

func (*ApiHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	postID, err := strconv.Atoi(strings.TrimPrefix(path, "/api/v1/post/"))
	if err != nil {
		common.Error(w, errors.New("请求路径无效, path:\t"+path))
		return
	}
	post, err := dao.GetPostByID(postID)
	if err != nil {
		log.Println(err)
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*ApiHandler) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	log.Println(token)
	log.Println(claims)
	userID := claims.Uid
	if err != nil {
		common.Error(w, errors.New("登陆已过期"))
		log.Println("登陆已过期")
		return
	}

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cid, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		post := &models.Post{
			Pid:        dao.CountGetAllPost() + 1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cid,
			UserId:     userID,
			ViewCount:  0,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cid, _ := params["categoryId"].(int)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		pid := int(params["pid"].(float64))
		viewCount := int(params["viewCount"].(float64))
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cid,
			UserId:     userID,
			ViewCount:  viewCount,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		log.Println(post.Title)
		service.UpdatePost(post)
		common.Success(w, post)
	}
}
func (*ApiHandler) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResponse := service.SearchPost(condition)
	common.Success(w, searchResponse)
}
