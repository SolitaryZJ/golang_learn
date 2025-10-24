package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/task/gorm/topic5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(139.224.237.37:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var dbx *sqlx.DB

func InitSqlxDB() *sqlx.DB {
	dsn := "root:123456@tcp(139.224.237.37:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dbx, err = sqlx.Connect("mysql", dsn) // 会真正 ping 一下
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	dbx.SetMaxOpenConns(100)
	return dbx
}
func main() {
	//db := InitDB()
	// SQL语句练习
	//topic1.Run(db)
	//topic2.Run(db)

	// sqlx
	//InitSqlxDB()
	//topic3.Run(dbx)
	//topic4.Run(dbx)

	topic5.Run(InitDB())
}
