package gutils

import (
	"os"
	"log"
	"io"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Dir(dir string) string {
	if PathExists(dir) {
		return dir
	}
	//mk dir
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Println(err)
	}
	return dir
}

//copy dir or file
func CopyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}

	source_open, err := os.Open(source)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()

	dest_open, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}


