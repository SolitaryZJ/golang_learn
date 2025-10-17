package main

import "fmt"

var a int

func main() {
	a = 5
	fmt.Println(a)
	a = 6
	fmt.Println(a)
	a := 7
	fmt.Println(a)

	if true {
		a := 8
		fmt.Println(a)
	}
	fmt.Println(a)
}
