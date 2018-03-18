package gutils

import (
	"time"
	"strings"
)

const (
	tF_ToSecond        = "2006-01-02 15:04:05"
	tF_ToSecondForName = "20060102150405"

	TF_Sec="yyyy-MM-dd HH:mm:ss"
	TF_Day="yyyy-MM-dd"
)

func Format(time time.Time, fmt string) string {
	gfmt := ToGoTimeFmt(fmt)
	return time.Format(gfmt)
}

func FormatToSecond(time time.Time) string {
	return time.Format(tF_ToSecond)
}

func FormatToSecondForFileName(time time.Time) string {
	return time.Format(tF_ToSecondForName)
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

func Parse(t, fmt string) (time.Time, error) {
	gtf := ToGoTimeFmt(fmt)
	return time.Parse(gtf,t)
}
