// Go语言基础教程 - 3. 运算符和控制结构
//
// 介绍Go语言中的运算符和程序控制结构

package main

import "fmt"

func main() {
	// 算术运算符
	fmt.Println("=== 算术运算符 ===")
	a, b := 10, 3
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	// 比较运算符
	fmt.Println("\n=== 比较运算符 ===")
	fmt.Printf("%d > %d ? %t\n", a, b, a > b)
	fmt.Printf("%d < %d ? %t\n", a, b, a < b)
	fmt.Printf("%d >= %d ? %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d ? %t\n", a, b, a <= b)
	fmt.Printf("%d == %d ? %t\n", a, b, a == b)
	fmt.Printf("%d != %d ? %t\n", a, b, a != b)

	// 逻辑运算符
	fmt.Println("\n=== 逻辑运算符 ===")
	x, y := true, false
	fmt.Printf("%t && %t = %t\n", x, y, x && y)
	fmt.Printf("%t || %t = %t\n", x, y, x || y)
	fmt.Printf("!%t = %t\n", x, !x)

	// 条件语句 if-else
	fmt.Println("\n=== 条件语句 ===")
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// switch语句
	fmt.Println("\n=== Switch语句 ===")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("星期一")
	case "Tuesday":
		fmt.Println("星期二")
	case "Wednesday":
		fmt.Println("星期三")
	case "Thursday":
		fmt.Println("星期四")
	case "Friday":
		fmt.Println("星期五")
	case "Saturday", "Sunday": // 多个条件
		fmt.Println("周末")
	default:
		fmt.Println("未知日期")
	}

	// switch可以没有条件表达式，类似if-else链
	switch {
	case score >= 90:
		fmt.Println("等级: A")
	case score >= 80:
		fmt.Println("等级: B")
	case score >= 70:
		fmt.Println("等级: C")
	case score >= 60:
		fmt.Println("等级: D")
	default:
		fmt.Println("等级: F")
	}

	// 循环结构
	fmt.Println("\n=== For循环 ===")
	// 基本for循环
	fmt.Println("基本for循环:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("第%d次循环 ", i)
	}
	fmt.Println()

	// while形式的for循环
	fmt.Println("while形式的for循环:")
	count := 0
	for count < 3 {
		fmt.Printf("count = %d ", count)
		count++
	}
	fmt.Println()

	// 无限循环（需要break跳出）
	fmt.Println("无限循环:")
	n := 0
	for {
		fmt.Printf("n = %d ", n)
		n++
		if n > 2 {
			break // 跳出循环
		}
	}
	fmt.Println()

	// range循环（遍历数组、切片、map等）
	fmt.Println("range循环:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// 只获取索引
	fmt.Println("只获取索引:")
	for i := range numbers {
		fmt.Printf("索引: %d\n", i)
	}

	// 只获取值
	fmt.Println("只获取值:")
	for _, value := range numbers {
		fmt.Printf("值: %d\n", value)
	}

	// continue和break的使用
	fmt.Println("\n=== continue和break ===")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		if i > 7 {
			break // 大于7时跳出循环
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}