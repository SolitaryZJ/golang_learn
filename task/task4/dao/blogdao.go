package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

var DB *gorm.DB

func init() {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(db)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	DB = db
}
