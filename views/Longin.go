package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"log"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login

	err := loginTemplate.WriteData(w, config.Cfg.Viewer)
	if err != nil {
		log.Println(err)
		return
	}
}
