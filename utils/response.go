package utils

import "github.com/gogf/gf/g/net/ghttp"

// Response 封装请求返回体
type Response struct {
	Code int
	Data interface{}
	Msg  string
}

// 后台的Response，为了兼容
type AdminResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// ErrorReponse 错误返回体
func ErrorResponse(msg string) *Response {
	return &Response{Code: 500, Msg: msg}
}

// SuccessReponse 成功返回体
func SuccessResponse(msg string) *Response {
	return &Response{Code: 200, Msg: msg}
}

// successWithData 成功返回体
func SuccessWithData(msg string, data interface{}) *Response {
	return &Response{Code: 200, Msg: msg, Data: data}
}

func Error(r *ghttp.Request) {
	r.Response.WriteJson(&AdminResponse{Code: 500})
}

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(&AdminResponse{Code: 200, Data: data})
}
