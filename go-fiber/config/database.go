package config

import (
	"alist/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB // 先定义，后使用，初始化后值改变，可直接引用
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("alist.db"), &gorm.Config{}) // 更新db的值
	if err != nil {
		panic(err)
	}
	DB = db //由于先前已赋值，无法再次赋值，需要等号
	// 此步骤惯例在这里进行
	if err := DB.AutoMigrate(models.File{}); err != nil {
		panic(err)
	}

	fmt.Println("数据库初始化完成")
}
