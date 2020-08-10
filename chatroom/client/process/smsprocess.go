package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess)SendGroupMes(content string) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content //消息内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal err=", err)
	}
	mes.Data = string(data)
	//4.序列化 mes
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal err=", err)
	}
	//5.将序列化的mes 发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err)
		return
	}
	return
}