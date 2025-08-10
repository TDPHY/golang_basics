package main

import "fmt"

/*
[加一](https://leetcode-cn.com/problems/plus-one/)
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {
	// 从最后一位开始处理
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果没有进位，直接返回
		if digits[i] != 10 {
			return digits
		}

		// 有进位，当前位变为0，继续处理前一位
		digits[i] = 0
	}
	// 如果所有位都进位，则需要扩展数组
	digits = make([]int, len(digits)+1)
	digits[0] = 1

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    //输出：[1, 2, 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) //输出：[4, 3, 2, 2]
	fmt.Println(plusOne([]int{9}))          //输出：[1, 0]
}
