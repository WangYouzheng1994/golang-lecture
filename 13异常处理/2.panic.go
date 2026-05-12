package main

import "os"

// 使用panic终止程序
func main() {
	funcPanic()
}

// 使用panic关键字终止程序
func funcPanic() {
	// 读取配置文件中，结果路径错了
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}
