package message

type User struct {
	//注意加tag保证序列化反序列化成功
	UserId int		`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string	`json:"userName"`
	UserStatus int 	`json:"userStatus"`
}