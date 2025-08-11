package main

import "fmt"

/*
指针
1. 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
  - 考察点 ：指针的使用、值传递与引用传递的区别。
*/
// 引用传递示例（会修改原值）
func increaseByPointer(num *int) {
	*num += 10 // 修改原始值
}

// 值传递示例（不会修改原值）
func increaseByValue(num int) {
	num += 10 // 只修改副本
}

func main() {
	value := 5
	fmt.Printf("修改前的值: %d\n", value)

	increaseByPointer(&value)
	fmt.Printf("引用传递修改后的值: %d\n", value)

	increaseByValue(value)
	fmt.Printf("值传递修改后的值: %d\n", value)
}
