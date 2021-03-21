package utils

import "strings"

// 去除空格和换行
func TrimSpaceLine(trimString string) string {
	trimString = strings.Trim(trimString, "\r\n\t")
	trimString = strings.Replace(trimString, " ", "", -1)
	// 去除换行符
	trimString = strings.Replace(trimString, "\n", "", -1)
	return trimString
}
