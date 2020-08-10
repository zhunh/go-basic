package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登陆成功的界面
func ShowMenu()  {
	fmt.Println("--------恭喜XXX登陆成功-----")
	fmt.Println("\t\t1.显示在线用户列表")
	fmt.Println("\t\t2.发送消息")
	fmt.Println("\t\t3.信息列表")
	fmt.Println("\t\t4.退出系统")
	fmt.Println("---------请输入1234-------")

	var key int
	var content string
	//因为我们总会使用到SmsProcess, 因此将其定义在 switch 外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户")
		outputOnlineUser()
	case 2:
		fmt.Println("你想说什么：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统。")
		os.Exit(0)
	default:
		fmt.Println("你的输入有误，请重新输入..")
	}
}
//和服务器端保持通讯
func ProcessServerMes(conn net.Conn)  {
	//创建一个transfer实例
	tf := &utils.Transfer{
		Conn:conn,
	}
	for{
		fmt.Printf("客户端%s 正在等待读取服务器发送的数据.\n",conn.RemoteAddr())
		mes, err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg(),err=",err)
			return
		}
		//读取到数据，进入下一步逻辑
		switch mes.Type {
			case message.NotifyUserStatusMesType://有人上线
				//1.取出 NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				//2.把这个用户的信息，状态保存到客户map[int]User 中
				updateUserStatusMes(&notifyUserStatusMes)
			case message.SmsMesType://有人群发消息
				OutputGroupMes(&mes)
			default:
			fmt.Println("未知消息类型")
		}
		//fmt.Println("mes=%v",mes)
	}
}