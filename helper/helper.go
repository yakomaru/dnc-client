package helper

import (
	"io/ioutil"
	"strings"
)

var fileTypes = map[string]bool{
	"3gp":  true,
	"avi":  true,
	"mov":  true,
	"mp4":  true,
	"m4v":  true,
	"m4a":  true,
	"mp3":  true,
	"mkv":  true,
	"ogv":  true,
	"ogm":  true,
	"ogg":  true,
	"oga":  true,
	"webm": true,
	"wav":  true,
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func listRecursion(dir string, localDir string, fileObj map[string]bool) {
	files, err := ioutil.ReadDir(dir + localDir)
	Check(err)
	for _, file := range files {
		if file.IsDir() {
			listRecursion(dir, localDir+file.Name()+"/", fileObj)
		} else {
			strArr := strings.Split(file.Name(), ".")
			if fileTypes[strArr[len(strArr)-1]] {
				fileObj[localDir+file.Name()] = true
			}
		}
	}
}

func ListFiles(dir string) map[string]bool {
	fileObj := make(map[string]bool)
	localDir := ""
	listRecursion(dir, localDir, fileObj)
	return fileObj
}
