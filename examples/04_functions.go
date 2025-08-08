// Go语言基础教程 - 4. 函数
//
// 函数是Go程序的基本构建块，用于执行特定任务

package main

import (
	"fmt"
	"math"
)

// 1. 基本函数定义
// func 函数名(参数列表) 返回值类型 { 函数体 }
func add(x int, y int) int {
	return x + y
}

// 参数类型相同时可以简写
func multiply(x, y int) int {
	return x * y
}

// 多返回值函数
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return x / y, nil
}

// 命名返回值
func rectangleStats(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 自动返回已命名的返回值
}

// 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 递归函数示例
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("=== 基本函数调用 ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	result2 := multiply(4, 6)
	fmt.Printf("4 * 6 = %d\n", result2)

	fmt.Println("\n=== 多返回值函数 ===")
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quotient)
	}

	// 忽略某个返回值使用下划线_
	quotient2, _ := divide(15, 5)
	fmt.Printf("15 / 5 = %.2f\n", quotient2)

	fmt.Println("\n=== 命名返回值函数 ===")
	area, perimeter := rectangleStats(5.0, 3.0)
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", area, perimeter)

	fmt.Println("\n=== 可变参数函数 ===")
	fmt.Printf("sum(): %d\n", sum())
	fmt.Printf("sum(1, 2): %d\n", sum(1, 2))
	fmt.Printf("sum(1, 2, 3, 4, 5): %d\n", sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	// 使用...展开切片作为可变参数传递
	fmt.Printf("sum(nums...): %d\n", sum(nums...))

	fmt.Println("\n=== 递归函数 ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	fmt.Println("\n=== 函数作用域 ===")
	scopeExample()
}

// 局部变量和全局变量示例
var globalVar = "我是全局变量"

func scopeExample() {
	localVar := "我是局部变量"
	
	fmt.Printf("全局变量: %s\n", globalVar)
	fmt.Printf("局部变量: %s\n", localVar)
	
	// 同名时局部变量优先
	globalVar := "我是局部变量，遮盖了全局变量"
	fmt.Printf("遮盖后的变量: %s\n", globalVar)
	
	// 在代码块中的变量作用域
	if true {
		blockVar := "我是代码块内的变量"
		fmt.Printf("代码块变量: %s\n", blockVar)
	}
	// fmt.Printf("无法访问blockVar: %s\n", blockVar) // 这行会编译错误
}

// 函数类型和函数作为值
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

// 闭包示例
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}