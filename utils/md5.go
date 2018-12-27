package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/pibigstar/go-todo/config"
)

// Md5 md5加密
func Md5(str string) string {
	secret := config.ServerConfig.GroupCodeSecret
	byte := []byte(fmt.Sprintf("%s%s", secret, str))

	sum := md5.Sum(byte)

	md5 := fmt.Sprintf("%x", sum)

	return md5
}
