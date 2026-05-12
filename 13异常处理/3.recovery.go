package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	recovery()
}

// 使用 defer+recovery捕获panic，进行程序恢复
func recovery() {
	defer func() {
		// 这是一种简写。
		if err := recover(); err != nil {
			fmt.Println(err) // 捕获异常，打印错误信息
			// 打印错误的堆栈信息
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) // 肯定会有一个panic
}
