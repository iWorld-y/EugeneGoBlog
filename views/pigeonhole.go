package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeResponse := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeResponse)
}
