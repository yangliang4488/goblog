package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yangliang4488/goblog/app/http/controllers"
	"github.com/yangliang4488/goblog/app/http/middlewares"
)

var pc *controllers.PageController = new(controllers.PageController)
var ac *controllers.ArticlesController = new(controllers.ArticlesController)

func RegisterWebRoutes(router *mux.Router) {

	router.Use(middlewares.ForceHtml)

	// 静态页面
	router.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	router.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	router.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	// 文章
	router.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[1-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("GET").Name("articles.delete")
}