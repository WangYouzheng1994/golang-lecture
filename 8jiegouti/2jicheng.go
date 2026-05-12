package main

import "fmt"

// Student 定义Student的继承结构体
type Child struct {
	Student
	Name string
	Age  int
}

// 给Child结构体绑了一个方法
func (c Child) childfunc1() {
	fmt.Println("childfunc1", c.Name, c.Age)
}

// func main() {
func main2() {
	student := Student{
		Name: "小王",
		Age:  1,
	}

	child := Child{
		Name: "小小王",
		Age:  12,
	}
	fmt.Println(student)
	fmt.Println(child)
	student.PrintInfo()
	// 孩子继承方法
	child.PrintInfo()
	child.childfunc1()
}
