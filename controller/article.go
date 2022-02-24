package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ArticleRender(w http.ResponseWriter, r *http.Request) {
	view := HtmlRender("article")
	query := r.URL
	filePath, _ := url.QueryUnescape(query.String())
	data, err := ioutil.ReadFile("./" + filePath + ".md")
	if err != nil {
		fmt.Println(err)
	}
	view.Execute(w, string(data))
}
