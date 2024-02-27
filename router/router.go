package router

import (
	"EugeneGoBlog/config"
	"EugeneGoBlog/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string
	Desc  string
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

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	path := config.Cfg.System.CurrentDir
	home := path + "/template/home.html"
	footer := path + "/template/layout/footer.html"
	header := path + "/template/layout/header.html"
	personal := path + "/template/layout/personal.html"
	post_list := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"

	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post_list, pagination)
	if err != nil {
		log.Println("解析模板错误: ", err)
	}
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	hr := &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	if err := t.Execute(w, hr); err != nil {
		log.Println(err)
	}
}

func Router() {
	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
