package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login

	loginTemplate.WriteData(w, config.Cfg.Viewer)
}
