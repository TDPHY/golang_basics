// Go语言基础教程 - 6. 结构体和方法
//
// 结构体是将零个或多个任意类型的命名字段组合在一起的数据结构

package main

import (
	"fmt"
	"math"
)

// 1. 定义结构体
// 使用type和struct关键字定义结构体
type Person struct {
	Name    string
	Age     int
	Address string
}

// 嵌套结构体
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person
	JobTitle string
	Salary   float64
	Address  // 嵌入结构体（匿名字段）
}

// 带标签的结构体
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-"表示忽略该字段
}

// 2. 方法定义
// 方法是带有接收者的函数
type Rectangle struct {
	Width, Height float64
}

// 为Rectangle类型定义方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 使用指针接收者修改结构体
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 3. 另一个示例：圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {
	fmt.Println("=== 结构体基础 ===")
	
	// 1. 创建结构体实例
	// 方式1：按字段顺序指定值
	person1 := Person{"张三", 30, "北京市"}
	fmt.Printf("person1: %+v\n", person1)
	
	// 方式2：使用命名字段
	person2 := Person{
		Name:    "李四",
		Age:     25,
		Address: "上海市",
	}
	fmt.Printf("person2: %+v\n", person2)
	
	// 方式3：零值初始化
	var person3 Person
	fmt.Printf("person3: %+v\n", person3)
	
	// 访问和修改字段
	person3.Name = "王五"
	person3.Age = 35
	person3.Address = "广州市"
	fmt.Printf("修改后的person3: %+v\n", person3)
	
	fmt.Println("\n=== 嵌套结构体 ===")
	emp := Employee{
		Person: Person{
			Name:    "赵六",
			Age:     28,
			Address: "深圳市",
		},
		JobTitle: "软件工程师",
		Salary:   15000.0,
		Address: Address{
			Street: "科技园南路1001号",
			City:   "深圳",
			Zip:    "518000",
		},
	}
	
	fmt.Printf("员工信息: %+v\n", emp)
	// 访问嵌入结构体的字段
	fmt.Printf("员工姓名: %s\n", emp.Name)        // 直接访问
	fmt.Printf("员工城市: %s\n", emp.Address.City) // 通过嵌入结构体访问
	
	fmt.Println("\n=== 结构体标签 ===")
	user := User{
		ID:       1,
		Username: "go_user",
		Email:    "go@example.com",
		Password: "secret123",
	}
	fmt.Printf("用户信息: %+v\n", user)
	
	fmt.Println("\n=== 方法 ===")
	// 创建Rectangle实例
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("矩形: %+v\n", rect)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())
	
	// 使用指针接收者的方法
	fmt.Println("\n--- 指针接收者方法 ---")
	fmt.Printf("缩放前: %+v\n", rect)
	rect.Scale(2.0)
	fmt.Printf("放大2倍后: %+v\n", rect)
	
	// 创建Circle实例
	circle := Circle{Radius: 5}
	fmt.Printf("\n圆形: %+v\n", circle)
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Perimeter())
	
	// 使用指针接收者
	fmt.Println("\n--- 圆形缩放 ---")
	fmt.Printf("缩放前半径: %.2f\n", circle.Radius)
	circle.Scale(3.0)
	fmt.Printf("放大3倍后半径: %.2f\n", circle.Radius)
	
	fmt.Println("\n=== 值接收者 vs 指针接收者 ===")
	// 值接收者不会修改原始值
	rect2 := Rectangle{Width: 3, Height: 4}
	fmt.Printf("原始矩形: %+v\n", rect2)
	rect2.Scale(2) // 虽然是值接收者，但Go会自动取地址
	fmt.Printf("调用Scale后: %+v\n", rect2)
	
	// 显式使用指针
	rect3 := Rectangle{Width: 2, Height: 3}
	rect3Ptr := &rect3
	rect3Ptr.Scale(3)
	fmt.Printf("显式使用指针: %+v\n", rect3)
}