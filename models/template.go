package models

import (
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

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	if err := t.Execute(w, data); err != nil {
		w.Write([]byte("Error"))
	}
}

func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog

	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)

		home := templateDir + "/home.html"
		footer := templateDir + "/layout/footer.html"
		header := templateDir + "/layout/header.html"
		personal := templateDir + "/layout/personal.html"
		postList := templateDir + "/layout/post-list.html"
		pagination := templateDir + "/layout/pagination.html"

		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": Date})
		t, err := t.ParseFiles(templateDir+"/"+viewName, home, header, footer, personal, postList, pagination)
		if err != nil {
			log.Println("解析模板错误: ", err)
		}
		tbs = append(tbs, TemplateBlog{Template: t})
	}
	return tbs
}

func InitTemplate(templateDir string) HtmlTemplate {

	tp := readTemplate([]string{
		"index", "category", "custom", "detail",
		"login", "pigeonhole", "writing",
	}, templateDir)
	return HtmlTemplate{
		Index:      tp[0],
		Category:   tp[1],
		Custom:     tp[2],
		Detail:     tp[3],
		Login:      tp[4],
		Pigeonhole: tp[5],
		Writing:    tp[6],
	}

}
