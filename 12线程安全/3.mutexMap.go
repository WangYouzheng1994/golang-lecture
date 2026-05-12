package main

import (
	"fmt"
	"sync"
	"time"
)

// 给map加锁进行读写操作

var waitGroupMutex sync.WaitGroup
var mp = map[string]string{}
var mutextLock sync.Mutex

// 通过锁的方式读取map
func readerMutexMap() {
	for {
		mutextLock.Lock()
		fmt.Println(mp["time"])
		mutextLock.Unlock()
	}
	waitGroupMutex.Done()
}

// 通过锁的方式写入map;
func writerMutexMap() {
	for {
		mutextLock.Lock()
		mp["time"] = time.Now().Format("15:04:05")
		mutextLock.Unlock()
	}
	waitGroupMutex.Done()
}

func main() {
	waitGroupMutex.Add(2)
	go readerMutexMap()
	go writerMutexMap()
	waitGroupMutex.Wait()
}
