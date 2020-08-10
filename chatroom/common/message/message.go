package message

const(
	SmsMesType				= "SmsMes"
	LoginMesType			= "LoginMes"
	LoginResMesType			= "LoginResMes"
	RegisterMesType			= "RegisterMes"
	RegisterResMesType 		= "RegisterResMes"
	NotifyUserStatusMesType	= "NotifyUserStatusMes"
)

//定义几个用户状态的常量
const(
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息
}

//消息定义
type LoginMes struct {
	UserId int  	`json:"userId"`//用户id
	UserPwd string 	`json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}

type LoginResMes struct {
	Code int 		`json:"code"`//返回状态码 500 表示该用户未注册 300 表示登陆成功
	UsersId []int	//增加字段，保存用户id的切片
	Error string 	`json:"error"`//返回错误类型
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code int 		`json:"code"` //返回 400 表示该用户已经占用，200 表示注册成功
	Error string 	`json:"error"`//返回 错误 信息
}
//为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//增加一个SmsMes, 发送的消息
type SmsMes struct {
	Content string `json:"content"`
	User //匿名结构体， 继承
}

