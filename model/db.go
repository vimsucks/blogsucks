package model

import "github.com/jinzhu/gorm"

var db *gorm.DB

func InitDb(d *gorm.DB) {
	db = d
	db.AutoMigrate(&Post{})
}
