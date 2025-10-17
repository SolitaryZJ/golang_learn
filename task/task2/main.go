package main

import "fmt"

func main() {
	x := 1
	intPointer(&x)
	fmt.Println(x)
}
func intPointer(x *int) {
	*x += 10
}
