package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User)
var CurUser model.CurUser  //在用户登录成功后完成初始化
//方法，处理返回的Notify
func updateUserStatusMes(notifyUserStatusMes *message.NotifyUserStatusMes)  {
	//适当优化, 先判断原来是否存在
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok{//原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	//原来有，则只需更改状态
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

//在客户端显示当前在线用户
func outputOnlineUser()  {
	//遍历 onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers{
		fmt.Println("用户id：\t", id)
	}
}