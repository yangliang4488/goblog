package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"unicode/utf8"

	"github.com/yangliang4488/goblog/app/models/article"
	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/route"
	"gorm.io/gorm"
)

type ArticlesController struct{}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)
		// 3. 渲染模板，将所有文章的数据传输进去
		tmpl.Execute(w, articles)
	}
}

type ArticleFormatData struct {
	Title, Body string
	URL         string
	Errors      map[string]string
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeUrl := route.Name2URL("articles.store")
	data := ArticleFormatData{
		Title:  "",
		Body:   "",
		URL:    storeUrl,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}
	return errors
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors)==0{
		insertId,err:=saveArticleDB()
	}else{

	}

}
