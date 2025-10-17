package main

import (
	"fmt"
	"time"
)

func main() {
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	// 中断switch
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break
		}
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

	// for {
	// 中断select
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("进过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后")
	}
	// }

	// 不使用标记
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break
		}
	}

	// 使用标记
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outter
		}
	}
}
