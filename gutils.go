package gutils

import (
	"io/ioutil"
	"strings"
	"path"
	"log"
)

func init() {
	all = make(map[string](map[string]interface{}))

	files, err := ioutil.ReadDir(dir_conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	for index, file := range files {
		log.Println("read No.", index+1, " confige file name:", file.Name())
		switch strings.ToLower(path.Ext(file.Name())) {
		case json_file:
			readConfByJson(file.Name())
		default:
			readConfByJson(file.Name())
		}
	}
}
