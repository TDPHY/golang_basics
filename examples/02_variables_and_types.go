// Go语言基础教程 - 2. 变量和数据类型
//
// Go语言有多种数据类型，包括基本类型和复合类型

package main

import "fmt"

func main() {
	// 变量声明方式1: var 变量名 类型
	var name string = "Go语言"
	var age int = 12
	var height float64 = 1.85
	var isPopular bool = true

	fmt.Println("=== 显式声明变量 ===")
	fmt.Printf("name: %s, type: %T\n", name, name)
	fmt.Printf("age: %d, type: %T\n", age, age)
	fmt.Printf("height: %.2f, type: %T\n", height, height)
	fmt.Printf("isPopular: %t, type: %T\n", isPopular, isPopular)

	// 变量声明方式2: var 变量名 类型（零值初始化）
	var city string
	var score int
	var weight float64
	var hasJob bool

	fmt.Println("\n=== 零值初始化 ===")
	fmt.Printf("city: '%s', type: %T\n", city, city)
	fmt.Printf("score: %d, type: %T\n", score, score)
	fmt.Printf("weight: %.2f, type: %T\n", weight, weight)
	fmt.Printf("hasJob: %t, type: %T\n", hasJob, hasJob)

	// 变量声明方式3: var 变量名 = 值（类型推导）
	var country = "中国"
	var population = 1400000000
	var area = 960.5

	fmt.Println("\n=== 类型推导 ===")
	fmt.Printf("country: %s, type: %T\n", country, country)
	fmt.Printf("population: %d, type: %T\n", population, population)
	fmt.Printf("area: %.1f, type: %T\n", area, area)

	// 变量声明方式4: 短变量声明（最常用）:=
	language := "Golang"
	version := 1.20
	isFast := true

	fmt.Println("\n=== 短变量声明 ===")
	fmt.Printf("language: %s, type: %T\n", language, language)
	fmt.Printf("version: %.1f, type: %T\n", version, version)
	fmt.Printf("isFast: %t, type: %T\n", isFast, isFast)

	// 多变量声明
	var (
		firstName = "张"
		lastName  = "三"
		age2      = 25
	)

	fmt.Println("\n=== 多变量声明 ===")
	fmt.Printf("姓名: %s%s, 年龄: %d\n", firstName, lastName, age2)

	// 常量声明
	const pi = 3.14159
	const (
		Monday = iota // iota从0开始，每次递增1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println("\n=== 常量 ===")
	fmt.Printf("π的值约为: %.5f\n", pi)
	fmt.Printf("星期几的枚举值: Monday=%d, Tuesday=%d, Wednesday=%d\n", Monday, Tuesday, Wednesday)

	// 复合类型简介
	fmt.Println("\n=== 复合类型简介 ===")
	// 数组
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组: %v, type: %T\n", numbers, numbers)

	// 切片
	letters := []string{"a", "b", "c", "d"}
	fmt.Printf("切片: %v, type: %T\n", letters, letters)

	// map
	person := map[string]interface{}{
		"name": "张三",
		"age":  30,
		"city": "北京",
	}
	fmt.Printf("map: %v, type: %T\n", person, person)
}