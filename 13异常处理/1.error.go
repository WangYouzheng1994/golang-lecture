package main

import (
	"errors"
	"fmt"
)

/*
go的异常处理有三种方式:
1. 向上抛出
2. 使用panic终止程序
3. 使用defer配合recovery捕获panic，recovery必须在defer中
*/

func Parent() error {
	// 使用下划线，填充完善语法，但是不真正的获取值
	err, _ := method() // 遇到错误向上抛
	return err
}
func method() (error, int) {
	return errors.New("出错了"), 1
}

// 向上抛出
func main() {
	fmt.Println(Parent())
}
