package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 判断是否存在路径
func ExistPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

func GetArticleTitle(path string) [2]string {
	var tmpTitleArr []string
	reg := regexp.MustCompile(`\\`)
	path = reg.ReplaceAllString(path, "/")
	tmpTitleArr = strings.Split(path, "/")
	var result [2]string
	if len(tmpTitleArr) > 2 {
		result[0] = tmpTitleArr[len(tmpTitleArr)-1]
		result[1] = tmpTitleArr[len(tmpTitleArr)-2]
	} else {
		result[0] = tmpTitleArr[len(tmpTitleArr)-1]
		result[1] = "默认"
	}
	return result
}

func GetArticlePath(path string) string {
	reg := regexp.MustCompile(`\\`)
	path = reg.ReplaceAllString(path, "/")
	return path
}
