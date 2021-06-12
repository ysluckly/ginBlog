package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


//数据库配置
const (
	userName = "root"
	password = "8mmwan+YS"
	ip = "localhost"
	port = "3306"
	dbName = "blog"
)

//Db数据库连接池
var (
	DB *sql.DB
)

//注意方法名大写，就是public
func InitMySQLDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//defer DB.Close()
	// 设置数据库最大连接数
	DB.SetMaxOpenConns(100)
	//设置连接的最大可复用时间
	DB.SetConnMaxLifetime(10 * time.Second)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}
