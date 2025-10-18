package main

import (
	"fmt"
	"time"
)

func main() {
	printNum()

	tm := TM{
		task: func() {
			fmt.Println("task1")
			time.Sleep(100 * time.Millisecond)
		},
	}
	taskTrigger(tm)

	time.Sleep(1000 * time.Millisecond)
}

type TM struct {
	task func()
}

func (tm *TM) runTask() {
	start := time.Now()
	tm.task()
	defer func() {
		fmt.Println("耗时：", time.Since(start))
	}()
}

func taskTrigger(t TM) {
	go t.runTask()
}

func printNum() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("even:", i)
			}
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("odd:", i)
			}
		}
	}()
}
