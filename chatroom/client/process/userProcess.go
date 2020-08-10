package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

//写一个函数，完成登陆功能
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf("userId = %d, userPwd = %s\n",userId,userPwd)
	//return nil
	//1.连接到服务器
	conn, err := net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err, err=", err)
		return
	}
	defer conn.Close()
	//2.创建发送消息结构体
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marsh err=",err)
		return
	}
	mes.Data = string(data)

	//将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marsh err=",err)
		return
	}

	//先发送data的长度给服务器
	//先获取到 data 的长度-> 转成切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送消息长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//fmt.Printf("客户端发送数据长度=%d, 内容=%s",len(data),string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//接收服务器返回的loginResMes
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err =tf.ReadPkg()
	if err != nil{
		fmt.Println("utils.ReadPkg err=",err)
		return
	}
	//将mes的Data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//初始化curUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//可以显示当前在线用户列表，遍历loginResMes.usersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId{
			//如果想在线用户不显示自己
			if v == userId{
				continue
			}
			fmt.Println("用户id\t", v)
			//完成客户端onlineUsers 初始化
			user := &message.User{
				UserId:v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}

		fmt.Println("\n\n")
		//这里再启动一个协程和服务器保持通讯
		go ProcessServerMes(conn)

		//1.登陆成功，循环显示登陆成功的菜单
		for{
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess)Register(userId int, userPwd, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err, err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn 发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//4.将registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data 赋值给 mes.Data字段
	mes.Data = string(data)
	//6.将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7.用transfer 发送数据
	//创建一个transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送数据给服务器
	fmt.Println("发送的注册信息：",string(data))
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("tf.WritePkg err=",err)
		return
	}
	//读取返回信息
	mes, err =tf.ReadPkg()
	if err != nil{
		fmt.Println("tf.ReadPkg err=",err)
		return
	}
	if err != nil{
		fmt.Println("注册发送信息错误 err=", err)
		return
	}
	//将mes的Data部分反序列化 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	fmt.Println("registerResMes",registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功， 请重新登录")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}