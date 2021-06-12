package main

import (
	"ginBlog/model"
	"ginBlog/proxy"
)

func main(){

	// 初始化数据库
	model.InitGORMDB()
	// 初始化路由组件
	proxy.InitRouter()
	return
}
