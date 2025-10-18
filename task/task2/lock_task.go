package main

import (
	"fmt"
	"sync"
)

var sharingNum int

func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lock.Lock()
				sharingNum++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(sharingNum)
}
