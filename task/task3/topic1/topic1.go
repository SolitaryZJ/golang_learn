package topic1

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID    int `gorm:"primary_key;autoIncrement"`
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	db.Debug().Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	var student Student
	db.Debug().Where("age > ?", 18).Find(&student)
	fmt.Println(student)

	db.Debug().Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "四年级")

	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
