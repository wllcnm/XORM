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

	//查询 Query
	//results, err := engine.Query("select * from user")
	//println(results)
	//queryString, err := engine.QueryString("select * from user")
	//fmt.Println(queryString)
	//queryInterface, err := engine.QueryInterface("select * from user")
	//fmt.Println(queryInterface)

	//get只能查询单条条件
	//GET(&user)中的参数为传入接受的参数,user为接收的参数
	//查询 GET //select * form user limit 1
	user := User{}
	engine.Get(&user)
	fmt.Println(user)

	//指定条件查询
	user = User{UserName: "jojo"}
	engine.Where("user_name=?", user.Id).Desc("id").Get(&user)
	fmt.Println(user)

	//获取指定字段的值,需要传入与表关联的地址engine.Table("表关联结构").where("查询条件").GET(""接受参数)
	user = User{UserName: "jojo", Id: 100}
	var name string
	engine.Table(&user).Where("id=?", user.Id).Cols("passwd").Get(&name)
	println(name)

	//通过find查询,支持多条记录查询
	var users []User
	//查询的条件为字符串时要加单引号
	engine.Where("passwd='123456'").And("user_name='jojo'").Limit(10, 0).Find(&users)
	//通过forr打印,前面的参数为序号,后面的参数为具体的值value
	for i, u := range users {
		fmt.Println(i, u)
	}
	//fmt.Println(users)

	//通过Count打印条数
	user = User{UserName: "jojo"}
	count, err := engine.Count(&user)
	println(count)

	//遍历查询
	//通过Rows查询
	rows, err := engine.Rows(&User{UserName: "jojo"})
	defer rows.Close()
	//bean为接受的类
	bean := new(User)
	for rows.Next() {
		//Scan就是将查询到的记录赋值给bean
		rows.Scan(bean)
		if bean.Id >= 1000 {
			fmt.Println(bean)
		}
	}
	//通过Iterate查询
	err = engine.Iterate(&User{UserName: "jojo"}, func(idx int, bean interface{}) error {
		newuser := bean.(*User)
		fmt.Println(newuser)
		return nil
	})
}
