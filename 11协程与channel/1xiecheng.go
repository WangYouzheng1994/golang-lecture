package main

import (
	"fmt"
	"time"
)

// 关键点：主线程结束后协程即使没有结束也会停止
// 关键字：go，用go来启动一个协程
func sing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
}

func main() {
	go sing()
	go sing()
	go sing()
	go sing()
	time.Sleep(2 * time.Second)
}
