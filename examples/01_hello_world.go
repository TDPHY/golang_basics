// Go语言基础教程 - 1. Hello World和基本结构
//
// Go（又称Golang）是Google开发的一种静态强类型编译型语言，
// 具有简洁、高效、可靠的特点。

// 每个Go程序都需要main包和main函数
package main

// 导入fmt包，用于格式化输入输出
import "fmt"

// main函数是程序的入口点
func main() {
	// 使用fmt.Println打印Hello World到控制台
	fmt.Println("Hello, 世界! 🌍")
	fmt.Println("Welcome to Go programming language tutorial!")
	
	// Go语言支持多种字符集，包括中文
	你好世界()
}

// 自定义函数示例
func 你好世界() {
	fmt.Println("This function name is written in Chinese!")
}