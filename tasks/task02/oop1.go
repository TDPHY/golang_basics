/*
面向对象
1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
  - 考察点 ：接口的定义与实现、面向对象编程风格。
*/
package main

import (
	"fmt"
)

// Shape 接口定义
type Shape interface {
	Area()
	Perimeter()
}

// Rectangle 结构体
type Rectangle struct {
	name string
}

// Rectangle 结构体实现Shape接口的Area()方法
func (r Rectangle) Area() {
	fmt.Printf("Rectangle 结构体的Area()方法实现，名称：%s\n", r.name)
}

// Rectangle 结构体实现Shape接口的Perimeter()方法
func (r Rectangle) Perimeter() {
	fmt.Printf("Rectangle 结构体的Perimeter()方法实现，名称：%s\n", r.name)
}

// Circle 结构体
type Circle struct {
	name string
}

// Circle 结构体实现Shape接口的Area()方法
func (r Circle) Area() {
	fmt.Printf("Circle 结构体的Area()方法实现，名称：%s\n", r.name)
}

// Circle 结构体实现Shape接口的Perimeter()方法
func (r Circle) Perimeter() {
	fmt.Printf("Circle 结构体的Perimeter()方法实现，名称：%s\n", r.name)
}

func main() {
	rectangle := Rectangle{name: "矩形"}
	rectangle.Area()
	rectangle.Perimeter()

	circle := Circle{name: "圆形"}
	circle.Area()
	circle.Perimeter()
}
