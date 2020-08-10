package process

import "fmt"

//UserMgr 实例在服务器端有且只有一个
//在很多地方都会使用到， 因此将其定义为全局变量

var(
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init()  {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1034),
	}
}

//添加在线用户
func (this *UserMgr)AddOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId] = up
}

//删除用户
func (this *UserMgr)DeleteOnlineUser(userId int)  {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr)GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr)GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//从map 中取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok{//说明查找用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}