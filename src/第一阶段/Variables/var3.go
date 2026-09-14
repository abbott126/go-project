package main

import "fmt"

var appName = "Go 学习示例"

func main() {
	name := "小李"
	age := 18
	score := 100
	isPass := true

	var (
		city   = "北京"
		school = "清华大学"
	)

	//fmt.Println(name, age, score, isPass, city, school)
	fmt.Printf("应用的名称： %s\n", appName)
	fmt.Printf("姓名: %s, 年龄: %d  分数: %.1f, 是否及格: %t\n ", name, age, score, isPass)
	fmt.Printf("城市 %s ,学校 %s\n", city, school)
}
