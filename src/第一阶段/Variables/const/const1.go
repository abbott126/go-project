package main

import "fmt"

// 定义常量

const name = "abbott"

//const name = "hello"

// 定义多个常量

const (
	StatusOk       = 200
	StatusNotFound = 404
	StatusError    = 500
)

// 带数据类型
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
)

func main() {
	fmt.Print("name")
	//const name = "hello world"
	name := "hello "
	fmt.Println(name)
	fmt.Printf("服务的状态： %s\n", StatusOk)
	fmt.Printf("找不到页面： %s\n", StatusNotFound)
	fmt.Printf("服务请求错误：%s\n", StatusError)
}
