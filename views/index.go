package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"errors"
	"log"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("(*HTMLApi) Index 首页获取数据失败:\t", err)
		index.WriteError(w, errors.New("系统内部错误\n"))
	}
	index.WriteData(w, hr)

}
