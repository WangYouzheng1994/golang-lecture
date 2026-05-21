package main

import "fmt"

// 关键点1：可以扩容
// 关键点2：与数组不同，他是引用传递，详见append_m1()
// 关键点3：容器中的内容的类型必须和声明的时候一样固定下来
// 约等于java的ArrayList
func main() {
	fmt.Println("-----slice_m1()")
	slice_m1()
	fmt.Println("-----slice_m2()")
	slice_m2()
	fmt.Println("-----slice_m3()")
	slice_m3()
	fmt.Println("-----slice_m4()")
	slice_m4()
	fmt.Println("-----append_m1()")
	append_m1()
	fmt.Println("-----slice_m5()")
	slice_m5()
}

// 切片，声明空切片
func slice_m1() {
	var list []string
	// 增加元素
	list = append(list, "1")
	list = append(list, "2")
	fmt.Println(list)
	list[0] = "123"
	fmt.Println(list)
}

// 声明并定义
func slice_m2() {
	var list []string = []string{"1", "2", "3"}
	fmt.Println(list)

	// 定义空切片
	var nilList = []string{}
	fmt.Println(nilList)

}

// 短定义
func slice_m3() {
	list := []string{"1", "2", "3"}
	fmt.Println(list)

	// 定义空切片
	nilList := []string{}
	fmt.Println(nilList)
}

// 通过make函数创建指定长度、容量的切片
// 长度代表已经根据类型做了初始化，可以通过索引直接修改了
// 容量代表目前的总长度，可继续扩容。容量主要为了解决数组拷贝的性能问题（与java 容器初始化定义长度同理）
func slice_m4() {
	// 初始化一个int类型的slice,前两个元素有值，总长度为5
	list := make([]int, 2, 5)
	fmt.Println(list)
	list[0] = 999
	list[1] = 998
	// list[2] = 997 会报错~因为并没有初始化到索引2
	fmt.Println(list)
}

// 从数组切出切片
func slice_m5() {
	list := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("待切割的数组", list)
	// 切割
	slice := list[:]
	fmt.Println("数组转切片", slice)
	// 基于索引下标切割，包头不包尾
	fmt.Println("把数组通过1~3切割，得到：", list[1:3])
}

// 切片是引用传递的
func append_m1() {
	list := []int{1, 2, 3}
	fmt.Println("初始化：", list)
	fmt.Println("dosomething返回：", doSomething(list))
	fmt.Println("doSomethin后", list) // 原容器没变，出现bug了

	// 必须要这么写！！
	list = doSomething(list)
	fmt.Println()
	fmt.Println("doSomethin后，要接收值", list)
}

// 正确写法（行业标准，永远不翻车）
func doSomething(s []int) []int {
	s = append(s, 100)
	return s // 必须返回
}
