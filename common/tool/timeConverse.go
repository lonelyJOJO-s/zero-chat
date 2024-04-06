package tool

import "time"

func Time2Str(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Time2TimeStamp(t time.Time) int64 {
	return t.Unix()
}
