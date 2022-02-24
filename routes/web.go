package routes

import (
	"net/http"
	"typemico/controller"
)

func InitWebRoute() {
	http.HandleFunc("/", controller.IndexRender)
	http.HandleFunc("/articles/", controller.ArticleRender)
	http.HandleFunc("/update", controller.ReloadArticleConfig)
}
