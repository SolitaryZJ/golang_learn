package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者：不需要 wg.Add，因为 close(ch) 会让 for range 自然结束
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
