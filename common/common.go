package common

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template")
}
