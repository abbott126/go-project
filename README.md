# golang  学习

## 第 1 阶段
Go 基础语法

目标：能独立写小型命令行程序，理解 Go 的核心习惯。
核心内容：

### 环境搭建：安装 Go（官网 go.dev）、配置 PATH、VS Code + Go 插件。
### 工具链：go run、go build、go mod、go test、go fmt、go vet。
### 基础语法：变量、常量、基本类型（int、string、bool、float64）、数组 vs 切片（slice）、map、结构体（struct）。
### 控制流：if、for（唯一循环）、switch、range。
### 函数：多返回值、命名返回值、defer、错误处理（error 接口，不要用异常）。
### 指针：传值 vs 传引用，理解何时用指针。
### 包与模块：package、import、Go Modules。

## 第 2 阶段
Go 核心能力
（结构体 / 接口 / 错误处理 / 泛型 / 包）

目标：掌握 Go 的灵魂——并发，写出自然的 Go 代码。
核心内容：

方法与接口：组合优于继承，接口是隐式实现。
错误处理深入：自定义 error、errors.Is/errors.As、wrap。
并发基础：goroutine、channel（有缓冲/无缓冲）、select、sync 包（Mutex、WaitGroup、Once 等）。
Context：取消、超时、传递请求级数据。
测试：testing 包、表驱动测试、benchmark、race detector（go test -race）。
标准库重点：fmt、os、io、encoding/json、net/http、time。

## 第 3 阶段
Go 并发编程
（goroutine / channel / context / sync）

## 第 4 阶段
Go 系统编程
（Linux / 文件 / Socket / TCP / HTTP）

## 第 5 阶段
Go Web 开发
（Gin / RESTful API / JWT / Middleware）

## 第 6 阶段
数据库与中间件
（MySQL / PostgreSQL / Redis / MQ）

## 第 7 阶段
Go 微服务
（gRPC / 服务发现 / 配置 / 日志 / tracing）

## 第 8 阶段
Go + Kubernetes
（client-go / controller / Operator）

##第 9 阶段
云原生开发
（Controller / Operator / CRD / Webhook）

## 第 10 阶段
AIOps
（Go + K8s + Agent + MCP/API + AI）