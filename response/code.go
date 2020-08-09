package response

// BadRequest 错误的请求，一般是参数验证失败
const BadRequest int = 1000

// SystemError 系统异常
const SystemError int = 1001

// Message 错误信息
var Message = map[int]string{
	BadRequest:  "参数错误",
	SystemError: "系统异常",
}
