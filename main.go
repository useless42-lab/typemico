package main

import (
	// "os/exec"

	"net/http"
	"typemico/config"
	"typemico/controller"
	"typemico/routes"
)

func main() {
	config.ReadBaseConfig()
	routes.InitWebRoute()
	controller.GetArticles()
	config.ReadArticleConfig()

	http.ListenAndServe(":"+config.Conf.Port, nil)
}
