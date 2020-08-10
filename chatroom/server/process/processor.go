package process

import (
	"chatroom/server/utils"
	"chatroom/common/message"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	//测试能否接收到客户端发送的消息
	fmt.Println("客户端发送：",mes)
	//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		//创建一个UserProcess实例
		up := &UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册的逻辑
		up := &UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。")
	}
	return
}

func (this *Processor) Processor() (err error) {
	//循环读客户端发送的信息
	for{
		//创建一个Transfer 实例， 完成读包的任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务器也推出..")
				return err
			}else {
				fmt.Println("readPkg err=",err)
			}
		}
		//fmt.Println("mes=", mes)
		//这里根据mes.Type 不同 使用不同的处理函数
		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}