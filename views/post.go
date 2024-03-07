package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detailTemplate := common.Template.Detail
	path := r.URL.Path
	// 把请求参数转为 int 类型的 postID
	postID, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(path, "/p/"), ".html"))
	if err != nil {
		detailTemplate.WriteError(w, errors.New("请求路径错误"))
		return
	}

	postResponse, err := service.GetPostDetail(postID)
	if err != nil {
		detailTemplate.WriteError(w, errors.New("文章访问失败"))
		return
	}
	detailTemplate.WriteData(w, postResponse)
}
