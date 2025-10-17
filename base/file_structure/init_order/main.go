package main

import (
	"fmt"
	_ "github.com/test/init_order/pkg1"
)

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init")
}
func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
func main() {
	fmt.Println("main.main method invoked!")
}
