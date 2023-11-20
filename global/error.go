package global

type ResponseError struct {
	ErrorCode int
	ErrorMsg  string
}

type ResponseErrors struct {
	BusinessError ResponseError
	ValidateError ResponseError
	TokenError    ResponseError
}

var Errors = ResponseErrors{
	BusinessError: ResponseError{40000, "业务错误"},
	ValidateError: ResponseError{42200, "请求参数错误"},
	TokenError:    ResponseError{403, "登录授权失效"},
}
