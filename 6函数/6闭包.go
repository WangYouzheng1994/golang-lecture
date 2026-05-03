package main

import "time"

// 关键点：闭包就是造出一个「只有一个方法、自带隐藏私有变量」的极简轻量对象
// 和js的闭包一个道理。
func main() {
	println(bibao_f1(1)(123, 321))
}

func bibao_f1(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	var sum int
	return func(numList ...int) int {
		for _, i2 := range numList {
			sum += i2
		}
		return sum
	}
}
