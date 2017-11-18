package gutils

import (
	"github.com/aswjh/excel"
	"time"
	"runtime"
)

const (
	xlsx = 51
	xls  = 56
	html = 44
)

func convertFile(inputFile, outputFile string, fmt interface{}) bool {
	old:=runtime.GOMAXPROCS(1)
	options := excel.Option{"Visible": false, "DisplayAlerts": true}
	xl, _ := excel.Open(inputFile, options)
	defer xl.Quit()
	time.Sleep(3000000000)
	xl.SaveAs(outputFile, fmt)
	runtime.GOMAXPROCS(old)
	return true
}

//xlsFilePath,outXlsxPath must Absolute path
//convert xls to xlsx
func ToXlsx(xlsFilePath, outXlsxPath string) bool {
	isSuc := convertFile(xlsFilePath, outXlsxPath, xlsx)
	return isSuc
}

//xlsFilePath,outXlsxPath must Absolute path
//convert xlsx to xlsx
func ToXls(xlsFilePath, outXlsxPath string)bool{
	isSuc := convertFile(xlsFilePath, outXlsxPath, xls)
	return isSuc
}

//xlsFilePath,outXlsxPath must Absolute path
//convert xls or xlsx to html
//Equivalent to excel soft "Save As"
func ToHtml(xlsFilePath, outHtmlPath string) bool {
	isSuc := convertFile(xlsFilePath, outHtmlPath, html)
	return isSuc
}
