package middleware

import (
	"errors"
	"strings"

	"github.com/gogf/gf/g/net/ghttp"
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
	validate := utils.CheckJwtToken(todoToken)
	if !validate {
		return "", errors.New("token is expired")
	}
	return token, nil
}

// CheckToken 检查token是否有效
func CheckToken(r *ghttp.Request) {
	// 如果是静态文件则不做token验证
	if r.IsFileRequest() {
		return
	}
	_, err := GetOpenID(r)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
}
