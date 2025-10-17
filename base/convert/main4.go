package main

import "fmt"

func main() {
	var i interface{} = 3
	switch v := i.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}
}
