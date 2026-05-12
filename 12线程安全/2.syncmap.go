package main

import (
	"fmt"
	"sync"
	"time"
)

// map是不允许并发读写的。报错：concurrent map read and map write
/*
如何实现线程安全的map？
方案1：给读写动作加锁
方案2：使用sync.map

*/

var waitGroup sync.WaitGroup

// 初始sync.map
var sync_mp = sync.Map{}

func readerSyncMap() {
	for {

		fmt.Println(sync_mp.Load("time"))
	}
	waitGroup.Done()
}
func writerSyncMap() {
	for {
		sync_mp.Store("time", time.Now().Format("15:04:05"))
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(2)
	go readerSyncMap()
	go writerSyncMap()
	waitGroup.Wait()

}
