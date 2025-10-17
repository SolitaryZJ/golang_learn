package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(string) {
		fmt.Println("run goroutine in closure params")
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
}
