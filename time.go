package gutils

import (
	"time"
	"strings"
)

const (
	TF_ToSecond        = "2006-01-02 15:04:05"
	TF_ToSecondForName = "20060102150405"
)

func FormatToSecond(time time.Time)string{
	return time.Format(TF_ToSecond)
}

func FormatToSecondForFileName(time time.Time)string{
	return time.Format(TF_ToSecondForName)
}

func ToGoTimeFmt(f string) string {
	f = strings.Replace(f, "yyyy", "2006", 1)
	f = strings.Replace(f, "MM", "01", 1)
	f = strings.Replace(f, "dd", "02", 1)
	f = strings.Replace(f, "hh", "15", 1)
	f = strings.Replace(f, "mm", "04", 1)
	f = strings.Replace(f, "ss", "05", 1)
	return f
}





