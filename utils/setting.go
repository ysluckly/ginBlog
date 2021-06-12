package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string // gin 模式
	HttpPort string // 服务端口

	Db       string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPassWd string
	DbName   string
)

// init 初始化
func init() {
	// 加载配置
	file, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Println("config init error", err)
		return
	}
	LoadServer(file)
	LoadDatabase(file)
	fmt.Println("config init success", AppMode, HttpPort, Db, DbHost, DbPort, DbUser, DbPassWd, DbName)
}

// LoadServer 加载服务配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

// LoadDatebase 加载数据库配置
func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWd = file.Section("database").Key("DbPassWd").MustString("admin23S")
	DbName = file.Section("database").Key("DbName").MustString("test1")
}
