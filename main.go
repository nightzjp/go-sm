package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
	"yishi-ai.com/go-sm/constant"
	model "yishi-ai.com/go-sm/models"
	"yishi-ai.com/go-sm/routes"
)

func init() {
	// 数据库链接
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库初始化失败.")
		return
	}
	sqlDB, dbError := db.DB()
	if dbError != nil {
		return
	}
	// 数据库自定义参数
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	constant.DB = db
	// 表结构迁移
	me := constant.DB.AutoMigrate(
		&model.User{},
		&model.SystemInfo{},
		&model.HttpConfig{},
		&model.MqttConfig{},
		&model.NetWorkConfig{},
		&model.Device{},
		&model.Algorithm{},
	)
	if me != nil {
		return
	}
}

func main() {
	gin.SetMode(gin.DebugMode) // 开发环境
	//gin.SetMode(gin.ReleaseMode)  // 生产环境
	r := routes.Routers()
	r.MaxMultipartMemory = int64(8 * constant.MB)
	_ = r.Run(":8089")
}
