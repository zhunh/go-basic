package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)
//在服务器启动后，就初始化一个userDao实例
//把它做成一个全局的变量，在需要和redis交互时，就直接使用即可
var(
	MyUserDao *UserDao
)

//定义一个UserDao 结构体,完成对User 结构体的各种操作
type UserDao struct {
	pool *redis.Client
}
//使用工厂模式， 创建一个UserDao 的实例
func NewUserDao(pool *redis.Client) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1.根据用户id 返回一个 user实例 + err
func (this *UserDao) getUserById(userId int) (user *message.User, err error) {
	//从userDao的连接池中取出一个连接
	//client := this.pool
	//defer client.Conn().Close()
	//通过给定id 去redis 查询这个用户
	res, err := this.pool.HGet("users", strconv.Itoa(userId)).Result()
	if err != nil{
		fmt.Println("client.HGet=", err)
	}
	fmt.Println("redis res=", res)
	//将查询结果 序列化
	err = json.Unmarshal([]byte(res),&user)
	if err != nil{
		fmt.Println("json.Unmarshal err=", err)
	}
	return
}

//完成登录校验 Login
//1.用户名密码和id都正确，返回一个user实例
//2.不正确返回错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	//defer client.Close()
	user, err = this.getUserById(userId)
	if err != nil{
		err = ERROR_USER_NOTEXISTS
		return
	}
	//这时获取到用户
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}
//完成注册
func (this *UserDao) Register(user *message.User) (err error) {
	_, err = this.getUserById(user.UserId)
	if err == nil{
		err = ERROR_USER_EXISTS
		return
	}
	//id在redis没有，可以入库
	data, err := json.Marshal(user)
	if err != nil{
		return
	}
	//入库
	_, err = this.pool.HSet("users", user.UserId, string(data)).Result()
	if err != nil{
		fmt.Println("HSet err=",err)
		return
	}
	return
}