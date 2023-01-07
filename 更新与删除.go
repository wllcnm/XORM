package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //一定记得导入数据库驱动,否则运行时会报错
	"time"
	"xorm.io/xorm"
)

func main() {
	//数据库连接基本信息
	var (
		userName  string = "root"
		password  string = "123456789"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "test"
		charset   string = "utf8mb4"
	)
	//构建数据库连接信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)

	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		println("数据库连接失败")
	}
	//使用结构体首字母必须大写
	type User struct {
		Id       int64
		UserName string
		Passwd   string
		Create   time.Time `xorm:"created`
		Updated  time.Time `xorm:"updated`
	}
	//将ID为101的记录的userName修改为cxk
	user := User{UserName: "cxk"}
	_, err = engine.ID(101).Update(&user)
	if err != nil {
		return
	}
	//删除 id为102的记录
	user = User{}
	_, err = engine.ID(102).Delete(&user)
	if err != nil {
		println("删除失败")
	}

	//使用sql语句
	engine.Exec("update user set user_name=? where id=100", "cxknb")
}
