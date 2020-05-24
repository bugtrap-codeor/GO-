# 《Go 语言从入门到实战》

## Go 静态编译型语言

go run xx.go//运行 xx.go  
go build xx.go//构建 xx.go

### 软件开发的新挑战

1.多核硬件架构  
2.超大规模分布式计算集群  
3.Web 模式导致的前所未有的开发规模和更新速度

### ·初识 Go 语言

1.Go 是一个开源的编程语言，它能很容易的构造简单、可靠且高效的软件  
2.Go 是从 2007 年末由 Robert Griesemer（Google V8 JS Engine 开发者、Hot Spot 开发者）,Rob Pike（Unix 的早期开发者、UTF-8 创始人）,Ken Thompson（Unix 的创始人、C 语言创始人、1983 年获图灵奖）主持开发，后来还加入了 lan Lance Taylor,Russ Cox 等人，并最终于 2009 年 11 月开源，在 2012 年发布了 Go1 稳定版本。

### GO 语言有哪些特点

简单 go:25 个关键字、C:37 个关键字、C++11:84 个关键字  
高效：垃圾回收、指针  
生产力：**复合 VS**继承  
云计算语言：docker、k8s  
区块链语言：ethereum、HYPERLEDGER  
·运行效率高，开发高效，部署简单  
·语言层面支持并发，易于利用多核实现并发  
·内置 runtime（作用：性能监控，GC 等）  
·简单易学，丰富的标准库，强大的网络库  
·内置强大的工具（gofmt)，跨平台编译，内嵌 C 支持

### GO 语言有那些应用

·服务器编程，如处理日志、数据打包、虚拟机处理、文件系统等  
·分布式系统，数据库代理器，中间件等  
·网络编程，目前使用最多最广泛的一块，Web 应用(国内爱好者开发的 go web 框架 beego)、API 应用等  
·云平台，目前云平台再逐步采用 GO 实现

### ·Go 语言环境搭建和 IDE 的安装

<https://golang.org/doc/install>
vscode、Atom<https://atom.io/、>...

### ·Go 语言基础语法

### ·Go 语言数据类型

### ·Go 语言变量和常量

### ·Go 语言控制语流程

基本程序结构

> package main//包，表明代码所在的模块（包）  
> import "fmt"//引入代码依赖  
> //功能实现
> func main() {  
> &emsp;fmt.Println("Hello World")  
> }

### 应用程序入口

1.必须是 main 包：package main 2.必须是 main 方法：func main（） 3.文件名不一定是 main.go

### 退出返回值

#### 与其他主要编程语言的差异

·Go 中 main 函数不支持任何返回值  
·通过 os.Exit 来返回状态

### 获取命令行参数

·main 函数不支持传入参数  
&emsp;func main(<del>arg []string</del>)  
·在程序中直接通过 os.Args 获取命令行参数

&emsp;package main //包，表明代码所在的模块（包）

&emsp;import (  
&emsp;&emsp;"fmt"  
&emsp;&emsp;"os"  
&emsp;)//引入代码依赖

&emsp;//功能实现  
&emsp;func main() {  
&emsp;&emsp;if len(os.Args) > 1 {  
&emsp;&emsp;fmt.Println("Hello World", os.Args[1])  
&emsp;&emsp;}  
&emsp;}

### 编写测试程序

1.源码文件以\_test 结尾：xxx_test.go  
2.测试方法名以 Test(包外可以访问)开头：
func TestXXX（t \*testing.T）{...}

### 实现 Fibonacci 数列

###变量赋值
与其他主要编程语言的差异
·赋值可以进行自动类型推断
·在一个赋值语句中可以对多个变量进行同时赋值

    // tmp := a
    // a = b
    // b = tmp
    a, b = b, a//变量值交换

快速设置连续值

    package constant_test

    import "testing"

    const (
    	Monday = 1 + iota
    	Tuesday
    	Wednesday
    )

    const (
    	Readable = 1 << iota
    	Writable
    	Executable
    )

    func TestConstantTry(t *testing.T) {
    	t.Log(Monday, Tuesday)
    }

    func TestConstantTry1(t *testing.T) {
    	a := 1 //0001
    	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
    }

###基本数据类型
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //alias for uint8
rune // alias for int32,represents a Unicode code point
float32 float64
complex64 complex128 ###类型转化
与其他主要编程语言的差异
1.Go 语言不允许隐式类型转换 2.别名和原有类型也不能进行隐式类型转换

指针类型与其他主要编程语言的差异 1.不支持指针运算
2.string 是值类型，其默认的初始化值为空字符串，而不是 nil
