package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name  string
	Age   int
	Call  func() byte
	Map   map[string]string
	Ch    chan string
	Arr   [32]uint8
	Slice []interface{}
	Ptr   *int
	once  sync.Once
}
type Custom struct {
	field1, field2, field3 byte
}
type Other struct{}

type Person1 struct {
	Name  string            `json:"name" gorm:"column:<name>"`
	Age   int               `json:"age" gorm:"column:<name>"`
	Call  func()            `json:"-" gorm:"column:<name>"`
	Map   map[string]string `json:"map" gorm:"column:<name>"`
	Ch    chan string       `json:"-" gorm:"column:<name>"`
	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
	Ptr   *int              `json:"-"`
	O     Other             `json:"-"`
}

type Custom1 struct {
	int
	string
	Other string
}

type Gender string

func isMale(g *Gender) bool {
	return *g == Male
}
func isFemale(g *Gender) bool {
	return *g == Female
}

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func main() {
	var gender = Male
	fmt.Println(gender)
	fmt.Println(isMale(&gender))
	fmt.Println(isFemale(&gender))
}
