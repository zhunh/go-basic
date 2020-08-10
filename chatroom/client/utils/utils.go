package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//想想它应该有哪些字段
	Conn net.Conn
	Buf	[8064]byte //传输时使用的缓冲
}
//将读取客户端数据包封装为函数
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据..")
	//先读取客户端发来的数据包长度值
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil{
		//err = errors.New("read pkg header err")
		return
	}
	//根据buf[0:4]转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err !=nil{
		//err = errors.New("read pkg body err")
		return
	}

	//反序列化mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}
//
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//1.发送消息长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//2.发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
