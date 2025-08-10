package main

import "fmt"

/*
[删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)
题目：给你一个有序数组 `nums` ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
Tips：可以使用双指针法，一个慢指针 `i` 用于记录不重复元素的位置，一个快指针 `j` 用于遍历数组，当 `nums[i]` 与 `nums[j]` 不相等时，将 `nums[j]` 赋值给 `nums[i + 1]`，并将 `i` 后移一位。
*/
func removeDuplicates(nums []int) int {
	// 空数组或只有一个元素
	if len(nums) <= 1 {
		return len(nums)
	}

	// 慢指针i指向不重复元素的末尾位置
	i := 0

	// 快指针j遍历数组
	for j := 1; j < len(nums); j++ {
		// 如果当前元素和前一个元素不同，则将当前元素复制到i+1的位置，并移动i指针
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1

}

func main() {
	nums := []int{1, 1, 2}
	length := removeDuplicates(nums)
	fmt.Println(length, nums[:length])
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length = removeDuplicates(nums)
	fmt.Println(length, nums[:length])
}
