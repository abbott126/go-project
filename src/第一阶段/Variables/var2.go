package main

import "fmt"

var global = "我是全局变量"

func main() {
	local := "我是局部变量"
	fmt.Println(local, global)
	a()
}

func a() {
	fmt.Println(global)
}
