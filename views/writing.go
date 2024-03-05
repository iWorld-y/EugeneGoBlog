package views

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writingTamplate := common.Template.Writing

	writingResponse := service.Writing()

	writingTamplate.WriteData(w, writingResponse)
}
