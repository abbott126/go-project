package main

import "fmt"

// 定义变量
func main() {
	var a int = 100
	var b string = "hello"
	// 短变量声明
	c := 19
	d := "helloworld "
	fmt.Println(a, b)
	fmt.Println(c, d)
	// 零值
	var name string
	var e int
	var age int
	var isName bool = true

	fmt.Println(name, age, isName, e)

	// 声明多个变量

	var aa, bb, cc int = 1, 2, 3

	var (
		name01 = "李四"
		age01  = 18
		height = 168.9
	)
	fmt.Print(aa, bb, cc)
	fmt.Println(name01, age01, height)
}
