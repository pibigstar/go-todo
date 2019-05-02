package utils

import "time"

// 日期格式化：2006-01-02 15:04:05
func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// 日期格式化：2006-01-02
func DateFormat(time time.Time) string {
	return time.Format("2006-01-02")
}
