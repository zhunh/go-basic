package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMes(mes *message.Message)  {
	//显示
	//反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil{
		fmt.Println("OutputGroupMes json.Unmarshal err=", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id：\t%d 对大家说：\t%s", smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
