package gutils

import (
	"testing"
	"../gutils"
)

func Test_CopyFile(t *testing.T) {

	src := "C:\\Users\\doudou\\Desktop\\岗位职责.docx"
	dest := "d:/copy/"

	gutils.CopyFile(src, dest, "dd.docx")

}
