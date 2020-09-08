package model

import (
	"strings"
)

//去除首尾特殊字符，转为小写
func filterString(s string) string {
	return strings.Replace(strings.ToLower(strings.TrimSpace(s)), ";", "", 1)
}

//首字母大写
func formatString(str string) string {
	s := strings.Split(str, "_")
	var upperStr string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if j == 0 {
				upperStr += strings.ToUpper(string(s[i][j]))
				continue
			}
			upperStr += string(s[i][j])
		}
	}
	return upperStr
}
