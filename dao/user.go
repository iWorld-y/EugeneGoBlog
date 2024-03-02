package dao

import "log"

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
