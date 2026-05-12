package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义三个全局通道
var moneyChan1 = make(chan int)    // 声明并初始化一个长度为0的信道
var nameChan1 = make(chan string)  // 声明并初始化一个长度为0的信道
var doneChan = make(chan struct{}) // 声明一个用于关闭的信道

// 接收三个参数，其中waitGroup是指针传递
func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

// 协程Select case的使用
// 关键点1：select case只能用于channel，
// 关键点2：select case的触发是两个条件：1. 通道有数据；2。 通道被关闭（传递一个0(int)进来
// 关键点3：使用自执行函数，配合wait.wait + defer函数，解决select case多余的通道执行
// 关键点4：select case，是多路复用的，如果同时触发，是随机触发一个case。一轮select case只会触发一个.注意！！！-> 官方定义的就是随机，避免case饥饿问题
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	// 现在的模式，就是购物接力
	//shopping("张三")
	//shopping("王五")
	//shopping("李四")
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
	go send("张三", 2, &wait)
	go send("王五", 3, &wait)
	go send("李四", 5, &wait)

	// 单独起了一个协程，等待wait都完成。
	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		// 后进先出（先执行）
		defer close(doneChan)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				fmt.Println("chufa:", money)
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				// 通道关闭的时候退出
				return
			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)

}
