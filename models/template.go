package models

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[(index+1)%len(strs)]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2024-02-27 19:15:00")
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	if err := t.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

}
func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		if _, err2 := w.Write([]byte(err.Error())); err2 != nil {
			log.Println(err)
			log.Println(err2)
		}
	}
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog

	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)

		home := templateDir + "home.html"
		footer := templateDir + "layout/footer.html"
		header := templateDir + "layout/header.html"
		personal := templateDir + "layout/personal.html"
		postList := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"

		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, postList, pagination)
		if err != nil {
			log.Println("解析模板错误: ", err)
			return nil, err
		}
		tbs = append(tbs, TemplateBlog{Template: t})
	}
	fmt.Printf("\033[1;37;41m%s\033[0m\n", "Template 读取完毕. ")
	return tbs, nil
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {

	tp, err := readTemplate([]string{
		"index", "category", "custom", "detail",
		"login", "pigeonhole", "writing",
	}, templateDir)
	if err != nil {
		log.Println("InitTemplate:\t", err)
		var htmlTemplate HtmlTemplate
		return htmlTemplate, err
	}
	fmt.Printf("\033[1;37;41m%s\033[0m\n", "Template 初始化完毕. ")

	return HtmlTemplate{
		Index:      tp[0],
		Category:   tp[1],
		Custom:     tp[2],
		Detail:     tp[3],
		Login:      tp[4],
		Pigeonhole: tp[5],
		Writing:    tp[6],
	}, nil

}
