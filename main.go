package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	orm "go-admin/database"
	"github.com/lexkong/log"
)

// @title go-admin API
// @version 0.0.1
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 
// @license.name MIT
// @license.url https://github.com/wenjianzhang/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	gin.SetMode(gin.DebugMode)
	log.Debug(config.DatabaseConfig.Port)

	err := gorm.AutoMigrate(orm.Eloquent)
	if err != nil {
		log.Fatalf("数据库初始化失败 err: %v", err)
	}

	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库基础数据初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}
}