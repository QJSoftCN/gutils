package gutils

import (
	"os"
	"log"
	"io"
	"path/filepath"
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

//copy file to dest dir
//destFileName if "" use sourceFileName
func CopyFile(sourceFile, destDir, destFileName string) bool {
	if sourceFile == "" || destDir == "" {
		log.Println("source or dest is null")
		return false
	}

	source_open, err := os.Open(sourceFile)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()

	//valid dest dir or create
	dir := Dir(destDir)

	if len(destFileName) == 0 {
		destFileName = filepath.Base(sourceFile)
	}

	destFile := filepath.Join(dir, destFileName)

	dest_open, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer dest_open.Close()

	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}
