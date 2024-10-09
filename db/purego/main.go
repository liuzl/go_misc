package main

import (
	"database/sql"
	"log"

	"gorm.io/driver/sqlite" // GORM 的 SQLite 驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // 使用 modernc.org/sqlite 作为驱动
)

type Product struct {
	ID    uint
	Code  string
	Price uint
}

func main() {
	// 使用 database/sql 配置 modernc.org/sqlite 驱动
	sqlDB, err := sql.Open("sqlite", "file:test.db")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 使用 GORM 配置数据库
	db, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 可选：启用日志以便调试
	})
	if err != nil {
		log.Fatalf("failed to connect database using gorm: %v", err)
	}

	// 自动迁移
	db.AutoMigrate(&Product{})

	// 插入数据
	db.Create(&Product{Code: "D42", Price: 100})

	// 查询数据
	var product Product
	db.First(&product, 1) // 根据主键查询
	log.Println(product)
}
