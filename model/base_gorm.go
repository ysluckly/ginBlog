package model

import (
	"fmt"
	"strings"
	"time"

	"ginBlog/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // 必须全局
	err error
)

func InitGORMDB(){
	path := strings.Join([]string{utils.DbUser, ":", utils.DbPassWd, "@tcp(",utils.DbHost, ":", utils.DbPort, ")/", utils.DbName, "?charset=utf8"}, "")
	db, err = gorm.Open(utils.Db, path)
	if err!= nil{
		fmt.Println("GORM db init faild")
		panic(err)
	}
	//defer db.Close()
	fmt.Println("GORM db init success")

	// 禁止默认表名的复数形式
	db.SingularTable(true)

	// 设置连接池最大空闲连接数
	db.DB().SetMaxIdleConns(10)

	// 设置数据库最大连接数
	db.DB().SetMaxOpenConns(10)

	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//自动检查结构是否变化，变化则自动进行迁移(包含自动创建表结构)
	db.AutoMigrate(&User{},&Category{},&Article{},&Profile{},&Comment{})
}