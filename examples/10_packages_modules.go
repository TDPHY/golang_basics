// Go语言基础教程 - 10. 包和模块管理
//
// 包是组织Go代码的基本单位，模块是Go 1.11引入的依赖管理机制

package main

import (
	"fmt"
	"math/rand"
	"time"
	
	// 可以导入自定义包（如果存在）
	// 例如: "github.com/yourname/yourpackage"
)

// 1. 包的基本概念
// 每个Go文件都属于一个包
// 包名通常与目录名相同
// main包是特殊的，包含程序入口点

// 2. 导入包的方式
// 标准包导入
import "os"

// 多个包导入
import (
	"log"
	"strconv"
)

// 3. 包的可见性
// 以大写字母开头的标识符是导出的（公共的）
// 以小写字母开头的标识符是非导出的（私有的）

// 公共变量
var PublicVariable = "我是公共变量"

// 私有变量
var privateVariable = "我是私有变量"

// 公共函数
func PublicFunction() string {
	return "这是公共函数"
}

// 私有函数
func privateFunction() string {
	return "这是私有函数"
}

// 公共结构体
type PublicStruct struct {
	PublicField  string // 公共字段
	privateField string // 私有字段
}

// 公共结构体的方法
func (p PublicStruct) PublicMethod() string {
	return "公共方法: " + p.PublicField
}

func (p PublicStruct) privateMethod() string {
	return "私有方法: " + p.privateField
}

// 4. 初始化函数
// 每个包可以包含多个init函数
// init函数在包初始化时自动执行，不能被其他函数调用

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("包初始化函数执行")
}

func init() {
	fmt.Println("另一个初始化函数执行")
}

func main() {
	fmt.Println("=== 包和模块基础 ===")
	
	// 使用导入的包
	fmt.Println("随机数:", rand.Intn(100))
	
	// 使用标准库包
	s := strconv.Itoa(42)
	fmt.Println("数字转字符串:", s)
	
	// 使用自定义包元素
	fmt.Println(PublicVariable)
	fmt.Println(PublicFunction())
	
	// 创建结构体实例
	ps := PublicStruct{
		PublicField:  "公共字段值",
		privateField: "私有字段值",
	}
	
	fmt.Println(ps.PublicMethod())
	// fmt.Println(ps.privateMethod()) // 这行会编译错误，因为privateMethod未导出
	
	fmt.Println("\n=== 包的别名导入 ===")
	// 可以为导入的包设置别名
	importedRand := rand.Intn(10)
	fmt.Println("使用别名的随机数:", importedRand)
	
	fmt.Println("\n=== 点导入 ===")
	// 使用点导入可以直接使用包中的标识符，无需前缀
	// import . "math"
	// fmt.Println(Sqrt(16)) // 不需要写成math.Sqrt(16)
	
	fmt.Println("\n=== 空白标识符导入 ===")
	// 使用空白标识符导入包，只执行包的初始化，不直接使用包中的标识符
	// import _ "image/png" // 只执行png包的初始化，注册PNG解码器
	
	fmt.Println("\n=== 模块管理 ===")
	fmt.Println("Go模块是Go 1.11引入的依赖管理机制")
	fmt.Println("使用go mod init命令初始化模块")
	fmt.Println("使用go mod tidy整理依赖")
	fmt.Println("使用go get添加依赖")
	
	// 模拟模块功能示例
	fmt.Println("\n--- 模块功能演示 ---")
	moduleExample()
}

// 模块功能示例函数
func moduleExample() {
	fmt.Println("这是一个模块功能示例")
	
	// 模拟使用外部包
	// 在实际项目中，这里会使用go.mod文件中定义的依赖
	data := []int{5, 2, 8, 1, 9, 3}
	sortedData := sortData(data)
	fmt.Printf("排序前: %v\n", data)
	fmt.Printf("排序后: %v\n", sortedData)
}

// 简单排序函数示例
func sortData(data []int) []int {
	// 创建副本避免修改原数据
	result := make([]int, len(data))
	copy(result, data)
	
	// 简单冒泡排序
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	return result
}