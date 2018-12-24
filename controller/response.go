package controller

type Response struct {
	code int
	data interface{}
	msg  string
}

// ErrorReponse 错误返回体
func errorResponse(msg string) *Response {
	return &Response{code: 500, msg: msg}
}

// SuccessReponse 成功返回体
func successResponse(msg string) *Response {
	return &Response{code: 200, msg: msg}
}

// successWithData 成功返回体
func successWithData(msg string, data interface{}) *Response {
	return &Response{code: 200, msg: msg, data: data}
}
