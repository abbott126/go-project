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

### 命令核心作用最常用场景
* go run编译并立即运行开发调试
* go build编译生成可执行文件发布、部署
* go mod模块与依赖管理管理第三方包
* go test运行测试验证代码正确性
* go fmt格式化代码保持风格统一
* go vet静态代码检查发现潜在 bug

### go run 命令示例：

```shell
go run main.go
go run .                  # 运行当前目录下所有 Go 文件
go run ./cmd/server       # 运行指定目录的包
```

### go build 命令示例：
```shell
go build                  # 编译当前目录，生成以目录名命名的可执行文件
go build -o myapp         # 指定输出文件名
go build ./cmd/server     # 编译指定包
```

### go test 命令示例：
```shell
go test                   # 运行当前包的所有测试
go test ./...             # 运行所有子包的测试
go test -v                # 显示详细输出
go test -race             # 开启竞态检测（强烈推荐）
go test -cover            # 显示测试覆盖率
go test -bench=.          # 运行基准测试
go test -run TestXxx      # 只运行指定测试函数
```
### go fmt 命令示例：
```shell
go fmt ./...              # 格式化当前模块所有文件
go fmt main.go            # 格式化单个文件
```
### go mod 命令示例：
```shell
go mod init 模块名初始化一个新模块（生成 go.mod）
go mod tidy自动添加缺失依赖、删除无用依赖（最常用）
go mod download下载依赖到本地缓存
go mod vendor把依赖复制到 vendor/ 目录
go mod graph查看依赖关系图
go mod why 包名解释为什么需要某个依赖
```

## 基础语法：
> 变量、常量、基本类型（int、string、bool、float64）、数组 vs 切片（slice）、map、结构体（struct）。

### 变量
### 常量
### 基本类型（int、string、bool、float64）
### 数组 vs 切片（slice）
### map
### 结构体（struct）



## 控制流：if、for（唯一循环）、switch、range。



## 函数：多返回值、命名返回值、defer、错误处理（error 接口，不要用异常）。



## 指针：传值 vs 传引用，理解何时用指针。


## 包与模块：package、import、Go Modules。