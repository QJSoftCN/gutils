package gutils

import (
	"testing"
	"../gutils"
	"fmt"
)

func Test_CopyFile(t *testing.T) {

	src := "C:\\Users\\doudou\\Desktop\\岗位职责.docx"
	dest := "d:/copy/"

	gutils.CopyFile(src, dest, "dd.docx")

}

func Test_GetExt(t *testing.T){
	src := "C:\\Users\\doudou\\Desktop\\岗位职责.docx"

	fmt.Println(gutils.GetFileExt(src))
}

func Test_GetFileName(t *testing.T){
	src := "C:\\Users\\doudou\\Desktop\\岗位职责.docx"

	fmt.Println(gutils.GetFileName(src))
}


