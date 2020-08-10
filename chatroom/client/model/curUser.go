package model

import (
	"chatroom/common/message"
	"net"
)
//因为在客户端，很多地方会使用到 curUser ,将其作为一个全局变量
type CurUser struct {
	Conn net.Conn
	message.User
}

