package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id
	path := r.URL.Path
	// 把请求参数转为 int 类型的 cId
	cId, err := strconv.Atoi(strings.TrimPrefix(path, "/c/"))
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	if err := categoryTemplate.WriteData(w, categoryResponse); err != nil {
		log.Println(err)
	}
}
