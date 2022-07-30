package models

import (
	"github.com/J-Nokwal/ihar_backend/pkg/configs"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	configs.Connect()
	db = configs.GetDB()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Likes{})
	db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
	db.Model(&Likes{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
	db.Model(&Likes{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Likes{}).AddUniqueIndex("like_unique_index", "user_id", "post_id")
}
