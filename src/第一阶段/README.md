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
方式一：完整声明（使用 var ）

```Go
var 变量名 类型
var 变量名 类型 = 初始值
```
示例：
```Go
var name string          // 声明但不赋值（会使用零值）
var age int = 25         // 声明并赋值
var isStudent bool = true
```



方式二：类型推断（最常用）
```Go
var name = "张三"        // 编译器自动推断为 string
var age = 25             // 推断为 int
```
方式三：短变量声明（函数内部最推荐）

```Go
name := "张三"           // 只能在函数内部使用
age := 25
isOk := true
```
注意：:= 只能在函数内部使用，包级别变量必须用 var。


#### 定义零值
Go 中变量声明后如果没有赋值，会自动拥有零值：

| 类型                                 | 零值           |
| ------------------------------------ | -------------- |
| int、float                           | 0              |
| bool                                 | false          |
| string                               | ""（空字符串） |
| 指针、切片、map、channel、函数、接口 | nil            |

示例：

```go
var a int
var b string
var c bool
fmt.Println(a, b, c)   // 输出：0  false
```

#### 一次声明多个变量

```go
// 方式1
var a, b, c int = 1, 2, 3

// 方式2（推荐）
var (
    name   string = "李四"
    age    int    = 30
    height float64 = 175.5
)

// 短声明
x, y := 10, 20
```

#### 变量作用域

**全局变量**：在函数外声明，整个包内可见。

**局部变量**：在函数或代码块内声明，只在该作用域有效。

```go
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
```

#### 注意事项与最佳实践

**未使用的变量会编译报错**（Go 强制要求）

```go
a := 10
// 如果不使用 a，编译会失败
```

**尽量使用短声明 :=**

- 代码更简洁
- 只在函数内部使用

**变量命名规范**

- 使用驼峰命名（camelCase）
- 包外可见用首字母大写（导出）
- 包内使用首字母小写

**不要过度使用全局变量**

- 优先使用局部变量和函数参数

完整的案例：

```go
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

```



### 常量**（Constants）**

> 常量是编译时就确定、运行时不能修改的值。Go 的常量设计很简洁且强大。

### 1. 基本声明方式

```go
const 常量名 类型 = 值
const 常量名 = 值          // 类型可省略（编译器推断）
```

示例：

```go
const PI float64 = 3.1415926
const AppName = "Go学习"
const MaxRetry = 3
```

------

### 2. 一次声明多个常量

```go
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// 或者带类型
const (
    Monday, Tuesday, Wednesday = 1, 2, 3
)
```

### iota —— 常量计数器（非常常用）

> iota 是 Go 特有的常量生成器，在 const 声明块中从 0 开始自增。

```go
const (
    Sunday    = iota  // 0
    Monday            // 1
    Tuesday           // 2
    Wednesday         // 3
    Thursday          // 4
    Friday            // 5
    Saturday          // 6
)
```

#### 常见高级用法：

**跳过某个值**

```go
const (
    a = iota  // 0
    b         // 1
    _         // 跳过 2
    c         // 3
)
```

**位运算（权限、标志位）**



### 基本类型（int、string、bool、float64）

### 数组 vs 切片（slice）
### map
### 结构体（struct）

## 控制流：if、for（唯一循环）、switch、range。



## 函数：多返回值、命名返回值、defer、错误处理（error 接口，不要用异常）。



## 指针：传值 vs 传引用，理解何时用指针。


## 包与模块：package、import、Go Modules。