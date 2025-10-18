package main

import "fmt"

func main() {
	x := 1
	intPointer(&x)
	fmt.Println(x)

	intSlice := []int{1, 2, 3}
	intSlicePointer(&intSlice)
	fmt.Println(intSlice)
}
func intPointer(x *int) {
	*x += 10
}
func intSlicePointer(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}
