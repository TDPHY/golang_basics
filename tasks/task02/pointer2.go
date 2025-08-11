package main

import "fmt"

/*
指针
2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
  - 考察点 ：指针运算、切片操作。、
*/
func doubleSlice(slice *[]int) {
	s := *slice
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前的切片: %v\n", numbers)

	doubleSlice(&numbers)
	fmt.Printf("修改后的切片: %v\n", numbers)
}
