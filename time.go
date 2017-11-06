package gutils

import "time"

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






