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
	session := engine.NewSession()
	defer session.Close()

	//开启事务
	session.Begin()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			//处理事务回滚
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	user1 := User{UserName: "zlxxx", Passwd: "cxknb", Id: 90}
	if _, err := session.Insert(&user1); err != nil {
		panic(err) //如果有错误,抛出一个错误
	}

	user2 := User{UserName: "jojo2", Passwd: "123456", Id: 1001}
	if _, err := session.Where("id=100").Update(&user2); err != nil {
		panic(err)
	}

	if _, err := session.Exec("delete from user where id=?", 9999); err != nil {
		panic(err)
	}
}
