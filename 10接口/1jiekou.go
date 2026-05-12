package main

import "fmt"

// 关键点1：隐式接口实现（duck typing 鸭子类型），即：当实现了接口的所有类型后，就是implements了接口

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// 多态函数
func sing(obj Animal) {
	obj.sing()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// 全部实现完之后，chicken就不再是一只普通的鸡了

// Cat 需要全部实现这些接口
type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Println("cat 唱")
}

func (c Cat) jump() {
	fmt.Println("cat 跳")
}
func (c Cat) rap() {
	fmt.Println("cat rap")
}

func main() {
	var animal Animal

	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()
	cat := Cat{"cat"}
	cat.sing()

	fmt.Println("----多态----")
	sing(animal)
	sing(cat)
}
