package mi

import "fmt"

//写一个函数，完成登陆功能

func login(userId int, userPwd string) (err error) {
	fmt.Printf("userId = %d, userPwd = %s",userId,userPwd)
	return nil
}
