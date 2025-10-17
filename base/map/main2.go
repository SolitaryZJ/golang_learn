package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"]++
		}
	}()

	go func() {
		for {
			m["a"]++
			fmt.Println(m["a"])
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}
