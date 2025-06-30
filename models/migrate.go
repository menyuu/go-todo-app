package models

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		log.Fatal("マイグレーションに失敗しました:", err)
	}
}
