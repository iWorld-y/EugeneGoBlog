package models

import "time"

type User struct {
	Uid       int       `json:"uid"`
	UserName  string    `json:"userName"`
	Passwd    string    `json:"passwd"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
