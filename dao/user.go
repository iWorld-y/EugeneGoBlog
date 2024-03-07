package dao

import (
	"EugeneGoBlog/models"
	"log"
)

func GetUser(userName, passwd string) *models.User {
	if row := DB.QueryRow(
		"select * from goblog.blog_user where user_name=? and passwd=? limit 1",
		userName,
		passwd); row.Err() == nil {
		// 若无错误
		var user models.User
		if err := row.Scan(&user.Uid,
			&user.UserName,
			&user.Passwd,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt); err == nil {
			return &user
		} else {
			log.Println("用户信息读取异常\t", err)
			return nil
		}
	} else {
		log.Println("登陆信息异常\t", row.Err())
		return nil
	}
}
func GetUserNameByID(userID int) string {
	if row := DB.QueryRow("select user_name from goblog.blog_user where uid=?", userID); row.Err() == nil {
		var userName string
		_ = row.Scan(&userName)
		return userName
	} else {
		log.Println("用户 ID 读取失败:\t", row.Err())
	}
	return ""
}
