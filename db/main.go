package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID   uint
	Name string
}

func init() {
	var err error
	dns := fmt.Sprintf("%s?_foreign_keys=true", "test.db")
	DB, err = gorm.Open(sqlite.Open(dns), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	if DB.AutoMigrate(&User{}) != nil {
		fmt.Println("failed to migrate")
		return
	}
	DB.Create(&User{Name: "John"})
	fmt.Println("Hello, World!")
}
