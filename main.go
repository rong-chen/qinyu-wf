package main

import (
	"qingyu-wf/init/db"
	"qingyu-wf/init/router"
	"qingyu-wf/init/viper"
)

func main() {
	// 初始化配置信息
	dbConfig := viper.Init()
	// 初始化数据库
	db.Init(dbConfig)
	// 初始化表
	//db.InitAutoMigrate()
	// 初始化路由
	router.Init()
}
