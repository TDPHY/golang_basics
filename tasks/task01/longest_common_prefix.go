package main

import "fmt"

/*
[最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
考察：字符串处理、循环嵌套
*/
func longestCommonPrefix(strs []string) string {
	// 如果字符串数组为空，返回空字符串
	if len(strs) == 0 {
		return ""
	}

	firstString := strs[0]

	// 以第一个字符串为基准
	for i := 0; i < len(firstString); i++ {
		char := firstString[i]

		// 检查其他字符串在位置i是否相同
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				return firstString[:i]
			}
		}
	}

	return firstString

}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
