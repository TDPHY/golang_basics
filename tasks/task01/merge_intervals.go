package main

import (
	"fmt"
	"sort"
)

/*
[合并区间](https://leetcode.cn/problems/merge-intervals/)
以数组 `intervals` 表示若干个区间的集合，其中单个区间为 `intervals[i] = [starti, endi]` 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，将第一个区间加入
	merged := [][]int{intervals[0]}

	// 遍历剩余的区间
	for i := 1; i < len(intervals); i++ {
		// 获取当前区间和前一个区间
		current := intervals[i]
		prev := merged[len(merged)-1]

		// 如果当前区间的起始位置小于等于前一个区间的结束位置，则说明有重叠
		if current[0] <= prev[1] {
			// 合并区间，更新结束位置为两者的较大值
			merged[len(merged)-1][1] = max(current[1], prev[1])
		}
	}
	return merged
}

func main() {
	fmt.Println(mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(mergeIntervals([][]int{{1, 4}, {4, 5}}))
}
