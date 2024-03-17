package common

// 确定消息类型
const (
	LoginMessageType       = "LoginMessage"
	LoginResultMessageType = "LoginResultMessage"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"`
}

type LoginMessage struct {
	UserId  int    `json:"user_id"`
	UserPwd string `json:"user_pwd"`
}

const (
	SucessCode = 200
	FailCode   = 500
	MissCode   = 400
)

type LoginResultMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
