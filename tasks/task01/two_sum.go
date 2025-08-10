package main

import "fmt"

/*
[两数之和](https://leetcode-cn.com/problems/two-sum/)**
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
*/
func twoSum(nums []int, target int) []int {
	// 使用map存储已遍历的数值和对应的索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算需要找到的另一个数
		complement := target - num

		// 在map中查找这个数
		if index, ok := numMap[complement]; ok {
			// 找到目标数值，返回两个数的索引
			return []int{index, i}
		}

		// 将当前数值和索引存入map
		numMap[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
