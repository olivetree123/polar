package response

// Response 返回值结构
type Response struct {
	UserID       string      `json:"userID"`
	RequestID    string      `json:"requestID"`
	Data         interface{} `json:"data"`
	StatusCode   int         `json:"statusCode"`
	CommandCode  int         `json:"commandCode"`
	ErrorMessage string      `json:"errorMessage"`
}

// MakeResponse 格式化返回值
func MakeResponse(requestID string, commandCode int, statusCode int, data interface{}, errorMessage string) *Response {
	response := Response{
		RequestID:    requestID,
		Data:         data,
		StatusCode:   statusCode,
		CommandCode:  commandCode,
		ErrorMessage: errorMessage,
	}
	//result, err := json.Marshal(response)
	//if err != nil {
	//	logrus.Errorln(err)
	//	return nil
	//}
	return &response
}

// SuccessResponse 成功时返回值
func SuccessResponse(requestID string, commandCode int, data interface{}) *Response {
	return MakeResponse(requestID, commandCode, 0, data, "")
}

// BadRequestResponse 返回参数错误提示
func BadRequestResponse(requestID string, commandCode int) *Response {
	return MakeResponse(requestID, commandCode, 1000, nil, "参数错误")
}

// APIResponse 返回结构
func APIResponse(requestID string, commandCode int, data interface{}) *Response {
	return SuccessResponse(requestID, commandCode, data)
}

// ErrorResponse xxx
func ErrorResponse(requestID string, commandCode int, statusCode int) *Response {
	return MakeResponse(requestID, commandCode, statusCode, nil, Message[statusCode])
}

// NewPushData 新建推送数据
func NewPushData(userID string, clientID string, code int, data interface{}) Response {
	return Response{
		Data:        data,
		UserID:      userID,
		CommandCode: code,
	}
}
