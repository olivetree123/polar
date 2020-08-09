package commands

// Command 命令结构
type Command struct {
	Code         int
	Args         map[string]interface{}
	UserID       string // 发起请求的用户 ID
	RequestID    string // 由客户端添加，一来用于请求的返回，二来方便问题的定位
	ClientID     string // 用于表明请求发起者的身份
	ConnectionID string // 客户端在建立连接时指定，用于标识连接
}

// PushData 主动推送的数据格式
type PushData struct {
	Code     int         //  命令代码
	Data     interface{} // 推送的数据
	UserID   string      // 推送的目标用户
	ClientID string      // 用于表明请求发起者的身份
}

// NewPushData 新建推送数据
func NewPushData(userID string, clientID string, code int, data interface{}) PushData {
	return PushData{
		Code:     code,
		Data:     data,
		UserID:   userID,
		ClientID: clientID,
	}
}

//func NewCommand(code int, args []byte, requestID string, userID string) *Command {
//	command := Command{
//		Code:      code,
//		Args:      args,
//		UserID:    userID,
//		RequestID: requestID,
//	}
//	return &command
//}
//
//func NewCommand2(code int, args interface{}, requestID string, userID string, buf *bytes.Buffer) error {
//	argsBytes, err := json.Marshal(args)
//	if err != nil {
//		return err
//	}
//	command := NewCommand(code, argsBytes, requestID, userID)
//	r, err := json.Marshal(command)
//	if err != nil {
//		return err
//	}
//	if _, err = buf.Write(r); err != nil {
//		return err
//	}
//	return nil
//}
