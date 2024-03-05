package service

import (
	"EugeneGoBlog/dao"
	"EugeneGoBlog/models"
	"EugeneGoBlog/utils"
	"errors"
	"log"
)

func Login(userName, passwd string) (*models.LoginResponse, error) {
	passwd = utils.Md5Crypt(passwd, "Eugene")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		// 若登陆失败
		log.Println("登陆失败, 用户名: ", userName, "\t密码: ", passwd)
		return nil, errors.New("账号或密码不正确")
	}
	uid := user.Uid

	// 使用 jwt 生成 Token 令牌
	token, err := utils.Award(&uid)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Token 未能生成")
	}
	userInfo := models.UserInfo{
		Uid:      user.Uid,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}
	var loginRes = &models.LoginResponse{
		Token:    token,
		UserInfo: userInfo,
	}
	log.Println(loginRes)
	return loginRes, nil
}
