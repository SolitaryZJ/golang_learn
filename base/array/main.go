package main

import "fmt"

func main() {
	// 仅声明
	var a [5]int
	fmt.Println("a = ", a)

	var marr [2]map[string]string
	fmt.Println("marr = ", marr)
	// map的零值是nil，虽然打印出来是非空值，但真实的值是nil
	// marr[0]["test"] = "1"

	// 声明以及初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 类型推导声明方式
	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("c = ", c)

	d := [3]int{3, 2, 1}
	fmt.Println("d = ", d)

	// 使用 ... 代替数组长度
	autoLen := [...]string{"auto1", "auto2", "auto3"}
	fmt.Println("autoLen = ", autoLen)

	// 声明时初始化指定下标的元素值
	positionInit := [5]string{1: "position1", 3: "position3"}
	fmt.Println("positionInit = ", positionInit)

	// 初始化时，元素个数不能超过数组声明的长度
	//overLen := [2]int{1, 2, 3}
}
