package process

import (
	"chatroom/server/model"
	"chatroom/server/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int //标识conn 属于哪个user
}
//编写通知所有在线的用户
//userId 通知其它的在线用户，我上线了
func (this *UserProcess)NotifyOtherOnlineUser(userId int)  {
	//遍历 onlineUsers, 然后一个一个的发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers{
		if id == userId{
			continue
		}
		//开始通知 [单独写一个方法]
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess)NotifyMeOnline(userId int)  {
	//组装NotifyUserStatusMesType
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将NotifyUserStatusMesType 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//将序列化的data 赋值给mes.Data
	mes.Data = string(data)
	//对mes 序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//发送，创建Transfer 实例
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("NotifyMeOnline err=", err)
	}

}
//编写一个函数serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes 中取出mes.Data 并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.声明一个LoginResMes
	var loginResMes message.LoginResMes

	//登录逻辑
	//这里使用redis数据库完成验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil{
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else{
		loginResMes.Code = 200
		fmt.Println(user,"登录成功")
		//这里用户登录成功，就把该用户放入 userMgr 中
		//将 登录成功的用户id userId 赋值给 this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this) //这个this 指的是func (this *UserProcess) ServerProcessLogin的 this
		//通知其它在线用户，我上线了
		this.NotifyOtherOnlineUser(this.UserId)
		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
	}
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
	//	loginResMes.Code = 200
	//}else{
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册再使用"
	//}
	//3.序列化loginMes
	data, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data 赋值给resMes.Data
	resMes.Data = string(data)
	//5.序列化res.Mes
	data, err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送res.Mes 给客户端
	//使用了分层模式，这里先创建Transfer 实例，然后发送数据
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
//处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//从mes 中提取出registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//使用redis 数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误.."
		}
	}else {
		registerResMes.Code = 200
	}
	//将registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil{
		fmt.Println("json.Marshal fail", err)
		return
	}
	//data 赋值给resMes
	resMes.Data = string(data)

	//序列化 resMes, 准备发送
	data, err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送数据
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}