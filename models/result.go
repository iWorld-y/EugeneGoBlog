package models

type Result struct {
	Error string      `json:"error"`
	Date  interface{} `json:"date"`
	Code  int         `json:"code"`
}
