package util

import (
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"regexp"
)

func IsUserName(s string) bool {
	if ok := regexp.MustCompile(`^[a-z][a-z\d]{3,29}$`).MatchString(s); !ok {
		return false
	}
	return true
}

func IsPassword(s string) bool {
	complexity := 0
	if ok := regexp.MustCompile(`[\da-zA-Z!@#$%^&*]{8,30}`).MatchString(s); !ok {
		return false
	}
	if ok := regexp.MustCompile(`\d`).MatchString(s); ok {
		complexity++
	}
	if ok := regexp.MustCompile(`[a-z]`).MatchString(s); ok {
		complexity++
	}
	if ok := regexp.MustCompile(`[A-Z]`).MatchString(s); ok {
		complexity++
	}
	if ok := regexp.MustCompile(`[!@#$%^&*]`).MatchString(s); ok {
		complexity++
	}
	if complexity < 4 {
		return false
	}
	return true
}

func CreatePassword(password string) (hashPassword string) {
	p := []byte(password)
	h, _ := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	hashPassword = string(h)
	return hashPassword
}

func ComparePassword(hashPassword string, password string) bool {
	p := []byte(password)
	h := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
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

func DecryptCFBToStruct(x any, secretKey string) error {
	vPrt := reflect.ValueOf(x).Elem()
	v := reflect.ValueOf(vPrt.Interface())
	n := v.NumField()
	if v.Kind() == reflect.Struct {
		for i := 0; i < n; i++ {
			k := v.Field(i)
			if k.Type().Kind() == reflect.String {
				str := k.String()
				plaintext, err := DecryptCFB(str, secretKey)
				if err != nil {
					return err
				}
				vPrt.Field(i).SetString(plaintext)
			}
		}
	}
	return nil
}
