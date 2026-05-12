package main

import "fmt"

// 关键点1：基于一个已有的类型，创造了新的类型，没有任何关联关系

func main() {
	fmt.Println(SuccessCode.GetMsg())
	var i int
	fmt.Println(int(SuccessCode) == i) // 必须要转成原始类型才能判断
}

// 自定义一个新的类型
type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7 // 校验失败的错误
	ServiceErrCode Code = 8 // 服务错误
)

// 可以新绑定方法
func (c Code) GetMsg() string {
	// 可能会有更加响应码返回不同消息内容的要求，我们在这个函数里面去实现即可
	// 可能还会有国际化操作
	return "成功"
}
