package common

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
func Success(w http.ResponseWriter, date interface{}) {
	result := models.Result{
		Error: "",
		Date:  date,
		Code:  200,
	}
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resultJson); err != nil {
		log.Println("登陆失败:\t", err)
	}
}

func Error(w http.ResponseWriter, err error) {
	result := models.Result{
		Error: err.Error(),
		Code:  403,
	}
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resultJson); err != nil {
		log.Println("登陆失败:\t", err)
	}
}

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
