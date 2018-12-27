package middleware

import (
	"errors"
	"strings"

	"gitee.com/johng/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/utils"
)

// GetOpenID 获取OpenID
func GetOpenID(r *ghttp.Request) (string, error) {
	todoToken := r.Header.Get("todo-token")
	todoToken = strings.Trim(todoToken, "")
	if todoToken == "" {
		return "", errors.New("token is nil")
	}
	token, err := utils.GetOpenIDFromToken(todoToken)
	if err != nil {
		return "", err
	}
	validata := utils.CheckJwtToken(token)
	if !validata {
		return "", errors.New("token已过期")
	}
	return token, nil
}

// CheckToken 检查token是否有效
func CheckToken(r *ghttp.Request) {
	_, err := GetOpenID(r)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
}
