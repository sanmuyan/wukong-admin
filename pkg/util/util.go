package util

import (
	"regexp"
)

func IsUserName(s string) bool {
	if ok := regexp.MustCompile(`^[a-z][a-z\d]{3,29}$`).MatchString(s); !ok {
		return false
	}
	return true
}

func IsSubPath(parentPath, subPath string) bool {
	p := []byte(parentPath)
	s := []byte(subPath)
	if p[0] != 47 || s[0] != 47 || len(p) >= len(s) {
		return false
	}
	j := 0
	for index := range p {
		if len(s) == index {
			break
		}
		if p[index] == s[index] {
			j = index
		}
	}
	// 获取两个路径相等的部分
	// p 相等的部分以 / 结尾  例子 /api/ /api/user
	// s 相等的部分下一个元素等于 / 例子 /api /api/user
	if p[j] == 47 || s[j+1] == 47 {
		return true
	}
	return false
}
