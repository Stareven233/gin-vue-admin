package main

// todo 删删改改交作业
// todo 如果真的有时间再写cyakka的api并写页面
// 其中api几乎都跟后台的不同，感觉后台其实没什么屌用，还搞得exe很大
// todo 没开cors中间件也没有跨域问题
// gorm comment会占用size长度
// gorm association返回的对象只可读取一次，即count完再find就没数据了
// 介绍preload连续调用及效果，提及related已被移除

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
