package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Md5Crypt 使用给定的字符串生成 md5
// @params str 待加密的字符串
// @params salt interface{} 加密盐
// @return str 返回生成的 md5 码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {

	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
