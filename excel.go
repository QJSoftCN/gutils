package gutils

import (
	"github.com/aswjh/excel"
	"time"
)

const (
	xlsx = 51
	xls  = 56
	html = 44
)

func convertFile(inputFile, outputFile string, fmt interface{}) bool {
	options := excel.Option{"Visible": false, "DisplayAlerts": true}
	xl, _ := excel.Open(inputFile, options)
	defer xl.Quit()
	xl.SaveAs(outputFile, fmt)
	time.Sleep(time.Second)
	return true
}

func ToXlsx(xlsFilePath, outXlsxPath string) bool {

	isSuc := convertFile(xlsFilePath, outXlsxPath, xlsx)

	return isSuc
}

func ToHtml(xlsFilePath, outHtmlPath string) bool {

	isSuc := convertFile(xlsFilePath, outHtmlPath, html)
	return isSuc
}
