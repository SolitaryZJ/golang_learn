package main

import "fmt"

func main() {
	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}
}
