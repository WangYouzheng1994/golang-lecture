package main

import "fmt"

func SetAge(info Student, age int) {
	info.Age = age
}

func SetAge1(info *Student, age int) {
	info.Age = age
}

// 本案例
// func main() {
func main3() {
	s := Student{
		Name: "枫枫",
		Age:  21,
	}
	fmt.Println(s.Age) // 21
	// 值传递 无法修改s
	SetAge(s, 18)
	fmt.Println(s.Age) // 21
	// 值传递 可以修改s
	SetAge1(&s, 17)
	// 提取指针，内部接收指针，所以修改了原对象
	fmt.Println(s.Age) // 17
}
