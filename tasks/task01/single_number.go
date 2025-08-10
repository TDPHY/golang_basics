package main

import "fmt"

/*
[只出现一次的数字](https://leetcode.cn/problems/single-number/)
题目：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
Tips：可以使用 `for` 循环遍历数组，结合 `if` 条件判断和 `map` 数据结构来解决，例如通过 `map` 记录每个元素出现的次数，然后再遍历 `map` 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		fmt.Println(num)
		m[num]++
		fmt.Println(m)
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums)) // 1
}
