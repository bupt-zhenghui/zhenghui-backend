package utils

import (
	"io/ioutil"
	"log"
	"os"
	"time"
	"zhenghui-backend/biz/consts"
)

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	var result []string

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				return result, err
			}
			result = append(result, temp...)
		} else if fi.Name() != ".DS_Store" {
			result = append(result, fullname)
		}
	}
	return result, nil
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func GetFileModTime(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Format(consts.FormatTime)
	}
	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Format(consts.FormatTime)
	}
	return time.Unix(fi.ModTime().Unix(), 0).Format(consts.FormatTime)
}
