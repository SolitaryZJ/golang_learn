package main

import "fmt"

func main() {

	var ch = make(chan int, 6)
	for i := 0; i < 6; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
