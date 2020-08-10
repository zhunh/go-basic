package main

import "fmt"

type dbConn struct {
	Type     string
	User     string
	Password string
	Host     string
	Database string
}

func main() {
	viper, err := loadConfig()
	if err != nil {
		fmt.Println("加载配置出错。")
		return
	}

	/*	conn := &dbConn{
		viper.Get("db.type").(string),
		viper.Get("db.user").(string),
		viper.Get("db.password").(string),
		viper.Get("db.host").(string),
		viper.Get("db.database").(string),
	}*/
	var conn *dbConn
	//conn := viper.Get("db").(dbConn)
	//for k, v := range viper.GetStringMapString("db") {
	//	fmt.Println(k, v)
	//}
	//conn := viper.GetStringMapString("db").(dbConn)
	viper.Unmarshal(&conn)
	fmt.Println(conn)
}
