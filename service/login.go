package service

import (
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"errors"
	"log"
)

func Login(userName, passwd string) (*models.LoginResponse, error) {
	user := dao.GetUser(userName, passwd)
	if user == nil {
		// 若登陆失败
		log.Println("登陆失败, 用户名: ", userName, "\t密码: ", passwd)
		return nil, errors.New("账号或密码不正确")
	}
	var loginRes = &models.LoginResponse{}
	return loginRes, nil
}
