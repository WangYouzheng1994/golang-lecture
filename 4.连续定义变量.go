package main

import "fmt"

var (
	t1 = "t1"
	t2 = "t2"
)

func main() {
	fmt.Println(t1, t2)
	m1()
	m2()
	m3()
}

func m1() {
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
}

func m2() {
	var a, b, c string = "1", "2", "3"
	fmt.Println(a, b, c)
}

// 短定义
func m3() {
	a, b, c := "10", "20", "30"
	fmt.Println(a, b, c)
}
