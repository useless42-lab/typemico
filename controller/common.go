package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"typemico/config"
)

func HtmlRender(path string) *template.Template {
	header := "template/" + config.Conf.Theme + "/public/header.html"
	footer := "template/" + config.Conf.Theme + "/public/footer.html"
	view, err := template.ParseFiles("template/"+config.Conf.Theme+"/"+path+".html", header, footer)
	if err != nil {
		fmt.Println(err)
	}
	return view
}

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ReloadArticleConfig(w http.ResponseWriter, r *http.Request) {
	os.Remove("./database/article.json")
	os.RemoveAll("./articles")
	GetArticles()
	config.ReadArticleConfig()
	msg, _ := json.Marshal(JsonResult{Code: 200, Msg: "更新完成"})
	w.Write(msg)
}

func GetArticles() {
	cmd := exec.Command("git", "clone", config.Conf.ArticleUrl, "articles")
	cmd.Run()
}
