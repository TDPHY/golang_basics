// Go语言基础教程 - 7. 接口
//
// 接口是方法的集合，定义了对象的行为

package main

import (
	"fmt"
	"math"
)

// 1. 定义接口
// 接口定义了一组方法签名
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. 实现接口的结构体
type Rectangle struct {
	Width, Height float64
}

// Rectangle实现Shape接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

// Circle实现Shape接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 3. 另一个接口示例
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// 4. 空接口
// interface{} 或 any (Go 1.18+) 可以表示任何类型
func PrintValue(v interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// 5. 类型断言
func DescribeShape(s Shape) {
	fmt.Printf("形状的面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
	
	// 类型断言检查具体类型
	if rect, ok := s.(Rectangle); ok {
		fmt.Printf("这是一个矩形，宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	} else if circle, ok := s.(Circle); ok {
		fmt.Printf("这是一个圆形，半径: %.2f\n", circle.Radius)
	} else {
		fmt.Println("未知的形状类型")
	}
}

// 6. 类型开关
func DescribeShapeWithSwitch(s Shape) {
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("矩形 -> 面积: %.2f, 宽度: %.2f, 高度: %.2f\n", 
			shape.Area(), shape.Width, shape.Height)
	case Circle:
		fmt.Printf("圆形 -> 面积: %.2f, 半径: %.2f\n", 
			shape.Area(), shape.Radius)
	default:
		fmt.Printf("未知形状: %T\n", shape)
	}
}

func main() {
	fmt.Println("=== 接口基础 ===")
	
	// 创建实现接口的实例
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}
	
	// 通过接口变量调用方法
	var s Shape
	s = rect
	fmt.Printf("矩形面积: %.2f\n", s.Area())
	fmt.Printf("矩形周长: %.2f\n", s.Perimeter())
	
	s = circle
	fmt.Printf("圆形面积: %.2f\n", s.Area())
	fmt.Printf("圆形周长: %.2f\n", s.Perimeter())
	
	fmt.Println("\n=== 多态性 ===")
	// 接口数组
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 2},
		Rectangle{Width: 5, Height: 2},
		Circle{Radius: 3},
	}
	
	for _, shape := range shapes {
		DescribeShape(shape)
	}
	
	fmt.Println("\n=== 类型开关 ===")
	for _, shape := range shapes {
		DescribeShapeWithSwitch(shape)
	}
	
	fmt.Println("\n=== 空接口 ===")
	// 空接口可以存储任何类型的值
	values := []interface{}{
		42,
		"Hello",
		3.14,
		true,
		Rectangle{Width: 1, Height: 2},
	}
	
	for _, v := range values {
		PrintValue(v)
	}
	
	fmt.Println("\n=== 接口与指针 ===")
	// 注意值和指针接收者的区别
	rect2 := Rectangle{Width: 3, Height: 4}
	var shape Shape = rect2 // 值拷贝
	fmt.Printf("面积: %.2f\n", shape.Area())
	
	var shape2 Shape = &rect2 // 指针
	fmt.Printf("面积: %.2f\n", shape2.Area())
	
	fmt.Println("\n=== 接口的隐式实现 ===")
	// Go中接口是隐式实现的，不需要显式声明实现了哪个接口
	// 只要类型实现了接口中的所有方法，就自动实现了该接口
}