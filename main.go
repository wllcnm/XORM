package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

// xorm文档:https://gitea.com/xorm/xorm/src/branch/master/README_CN.md
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
	//插入对象
	user := User{Id: 100, UserName: "jojo", Passwd: "123456"}
	result, _ := engine.Insert(&user)
	println(result)
	if result >= 1 {
		println("插入成功")
	} else {
		println("插入失败")
	}

	//插入多条对象
	user1 := User{Id: 101, UserName: "jojo", Passwd: "123456"}
	user2 := User{Id: 102, UserName: "jojo", Passwd: "123456"}
	engine.Insert(&user1, &user2)
	println(result)
	if result >= 1 {
		println("插入成功")
	} else {
		println("插入失败")
	}

	//插入对象切片
	var users []User
	users = append(users, User{Id: 103, UserName: "jojo", Passwd: "123456"})
	users = append(users, User{Id: 104, UserName: "jojo", Passwd: "123456"})
	engine.Insert(&users)
	err = engine.Sync(new(User))
	if err != nil {
		println("表结构同步失败")
	}
}
