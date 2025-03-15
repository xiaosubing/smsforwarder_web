package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func NewDB() {
	// 1. 连接数据库（启用WAL模式提升并发性能）
	db, err := gorm.Open(sqlite.Open("/home/huihui/go/src/study/bg/test.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				//SlowThreshold: time.Second, // 慢查询阈值
				//LogLevel:      logger.Info, // 日志级别
				//Colorful:      true,        // 彩色日志
				LogLevel: logger.Silent,
			},
		),
	})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
		return
	}

	// 自动键表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	err = db.AutoMigrate(&Message{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	DB = db
}
