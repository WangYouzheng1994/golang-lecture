package main

import "fmt"

// 关键点：等价于static，按序执行，在main以前
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}
func main() {
	fmt.Println("main")
}
func init() {
	fmt.Println("init5")
}
