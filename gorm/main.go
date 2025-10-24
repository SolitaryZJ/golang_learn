package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/lesson2"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

func InitDB(dst ...interface{}) *gorm.DB {
	dsn := "root:123456@tcp(139.224.237.37:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(dst...)
	return db
}

func main() {
	//db, err := gorm.Open(mysql.Open("root:123456@tcp(139.224.237.37:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)

	db := InitDB(&Product{}, &Child{})
	//lesson1.Run(db)
	lesson2.Run(db)
}
