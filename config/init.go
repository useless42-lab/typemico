package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"time"
	"typemico/models"
	"typemico/utils"
)

const BaseConfigPath string = "./config.json"
const ArticleConfigPath string = "./database/article.json"

type Config struct {
	models.BaseConfig
	ArticleConfig []models.ArticleConfig
}

var Conf Config

func CheckBase() (bool, error) {
	r, err := utils.ExistPath(BaseConfigPath)
	return r, err
}

func CheckArticle() (bool, error) {
	r, err := utils.ExistPath(ArticleConfigPath)
	return r, err
}

func ReadBaseConfig() {
	isExist, err := CheckBase()
	if err != nil {
		fmt.Println(err)
	}
	if isExist {
		rf, err := ioutil.ReadFile(BaseConfigPath)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(rf, &Conf.BaseConfig)

	}
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getCreateTime(path string) time.Time {
	finfo, err := os.Stat(path)
	modTime := finfo.ModTime()
	if err != nil {
	}
	// return modTime.Format("2006-01-02 15:04:05")
	// return modTime.Format("2006-01-02")
	return modTime
}

func InitArticleConfig() []byte {
	files, err := utils.WalkDir("./articles", ".md")
	if err != nil {
		fmt.Println(err)
	}

	var articleItem []models.ArticleConfig
	for i := 0; i < len(files); i++ {
		tmpArticleItem := models.ArticleConfig{
			Title:    strings.Replace(utils.GetArticleTitle(files[i])[0], ".md", "", 1),
			Category: utils.GetArticleTitle(files[i])[1],
			Path:     strings.Replace(utils.GetArticlePath(files[i]), ".md", "", 1),
			// Date:     getCreateTime(files[i]).Format("2006-01-02"),
		}
		articleItem = append(articleItem, tmpArticleItem)
	}
	data, err := json.Marshal(articleItem)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func ReadArticleConfig() {
	isExist, err := CheckArticle()
	if err != nil {
		fmt.Println(err)
	}
	if isExist {
		rf, err := ioutil.ReadFile(ArticleConfigPath)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(rf, &Conf.ArticleConfig)

	} else {
		os.Mkdir("./database", os.ModePerm)
		fp, err := os.OpenFile(ArticleConfigPath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err)
		}
		defer fp.Close()
		_, err = fp.Write(InitArticleConfig())
		json.Unmarshal(InitArticleConfig(), &Conf.ArticleConfig)
	}
}
