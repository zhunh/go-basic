package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//暂时不需要字段
}

//转发消息
func (this *SmsProcess)SendGroupMes(mes *message.Message)  {
	//遍历服务器端的 onlineUsers map[int]*UserProcess
	//将消息发出
	//取出mes 的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil{
		fmt.Println("SendGroupMes json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers{
		//这里过滤掉自己，不发给自己
		if id == smsMes.UserId{
			continue
		}
		this.SendMsg(data, up.Conn)
	}
}

func (this *SmsProcess)SendMsg(data []byte, conn net.Conn)  {
	//创建 transfer
	tf := &utils.Transfer{
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("服务器转发消息失败 err=", err)
	}
}