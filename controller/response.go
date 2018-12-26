package controller

// Response 封装请求返回体
type Response struct {
	Code int
	Data interface{}
	Msg  string
}

// ErrorReponse 错误返回体
func errorResponse(msg string) *Response {
	return &Response{Code: 500, Msg: msg}
}

// SuccessReponse 成功返回体
func successResponse(msg string) *Response {
	return &Response{Code: 200, Msg: msg}
}

// successWithData 成功返回体
func successWithData(msg string, data interface{}) *Response {
	return &Response{Code: 200, Msg: msg, Data: data}
}
