# Go 基础语法

> 目标：能独立写小型命令行程序，理解 Go 的核心习惯。
核心内容：

## 环境搭建：安装 Go（官网 go.dev）、配置 PATH、VS Code + Go 插件。


### 检查环境

```bash
 [main] go version 
go version go1.24.5 darwin/arm64

```
### 环境安装完成


### 运行第一个hello world的程序

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
运行代码：
```bash
go run helloworld.go
hello world
```



## 工具链：go run、go build、go mod、go test、go fmt、go vet。



## 基础语法：变量、常量、基本类型（int、string、bool、float64）、数组 vs 切片（slice）、map、结构体（struct）。




## 控制流：if、for（唯一循环）、switch、range。



## 函数：多返回值、命名返回值、defer、错误处理（error 接口，不要用异常）。



## 指针：传值 vs 传引用，理解何时用指针。


## 包与模块：package、import、Go Modules。