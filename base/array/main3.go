package main

import "fmt"

func main() {
	// 三维数组
	a := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}

	layer1 := a[0]
	layer2 := a[0][1]
	element := a[0][1][1]
	fmt.Println(layer1)
	fmt.Println(layer2)
	fmt.Println(element)

	// 多维数组遍历时，需要使用嵌套for循环遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
		for j, inner := range v {
			fmt.Println("inner, index = ", j, "value = ", inner)
		}
	}
}
