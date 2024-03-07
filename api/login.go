package api

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"net/http"
)

func (*ApiHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 接受用户名与密码
	// 返回数据
	param := common.GetRequestJsonParam(r)
	userName := param["username"].(string)
	passwd := param["passwd"].(string)
	if loginResponse, err := service.Login(userName, passwd); err == nil {
		common.Success(w, loginResponse)
	} else {
		common.Error(w, err)
		return
	}
}
