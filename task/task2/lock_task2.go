package main

import (
	"sync"
	"sync/atomic"
)

var num atomic.Int32

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				num.Add(1)
			}
		}()
	}
	wg.Wait()
	println(num.Load())
}
