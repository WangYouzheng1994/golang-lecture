package main

import (
	"fmt"
	"sync"
	"time"
)

// 关键点1：Javaer，100%理解countdownlatch即可。
// Add(n int)：设置需要等待的协程数量（必须在启动协程前调用）
// Done()：等价于 Add(-1)，协程完成任务时调用
// Wait()：阻塞主线程 / 主协程，直到计数器变为 0（所有协程都调用了 Done）

func main() {
	// 声明有4个线程
	wait.Add(4)
	go testsing()
	go testsing()
	go testsing()
	go testsing()
	// 阻塞等待
	wait.Wait()
	fmt.Println("主线程结束")
}

var (
	wait = sync.WaitGroup{}
)

func testsing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	// 通知完成了一个~
	wait.Done()
}
