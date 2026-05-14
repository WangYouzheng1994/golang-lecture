package main

import (
	"fmt"
	"time"
)

/*
关键点1：通道操作符 <-；从右往左读：
1. c <- chanA 从通道chanA中读取数据给c
2. chanA <- c ，把c的数据写入到通道chanA中
*/

/*
关键点2： chan代表双向通道；
1. chan<-代表只写通道；
2. <-chan 代表只读通道
*/
func main() {
	chan_m1()
	chan_m2()
}

// 通道的基本使用
func chan_m1() {
	var c chan int // 声明一个传递整形的通道
	// 初始化通道
	c = make(chan int, 1) //  初始化一个 有一个缓冲位的通道
	// 写入通道
	c <- 1
	//c <- 2           // 会报错 deadlock，因为只有一个长度位
	fmt.Println(<-c) // 取值
	//fmt.Println(<-c) // 再取也会报错  deadlock

	// 写入通道
	c <- 2
	// 取值 + 通道状态
	n, ok := <-c
	fmt.Println(n, ok)
	close(c) // 关闭协程
	//c <- 3   // 关闭之后就不能再写或读了  send on closed channel
	fmt.Println(c)
}

// make生成无缓冲的通道
func chan_m2() {
	// 无缓冲通道
	ch := make(chan int)
	go func() {
		// 阻塞5秒（time.after返回的是通道。我们需要用<-接收，来进行阻塞
		<-time.After(5 * time.Second)
		ch <- 666
		// 阻塞等待接收
	}()
	fmt.Println("从无缓冲通道中读取，", <-ch)

}
