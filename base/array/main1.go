package main

import "fmt"

func main() {
	a := [5]int{5, 4, 3, 2, 1}

	// 方式1，使用下标读取数据
	element := a[2]
	fmt.Println("element = ", element)

	// 方式2，使用range遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
	}

	for i := range a {
		fmt.Println("only index, index = ", i)
	}

	// 读取数组长度
	fmt.Println("len(a) = ", len(a))
	// 使用下标，for循环遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println("use len(), index = ", i, "value = ", a[i])
	}
}
