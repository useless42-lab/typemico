package controller

import (
	"net/http"
	"typemico/config"
)

type ReponseData struct {
	Config config.Config
}

func IndexRender(w http.ResponseWriter, r *http.Request) {
	rd := ReponseData{
		Config: config.Conf,
	}
	view := HtmlRender("index")
	view.Execute(w, rd)
}
