package main

import "fmt"

/*
[回文数](https://leetcode-cn.com/problems/palindrome-number/)
题目：给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
Tips：回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例：121 是回文，而 123 不是。
*/

func isPalindrome(x int) bool {
	// 负数和末尾0直接返回false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 反转数字
	var revertedNum int = 0
	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		fmt.Println(revertedNum)
		x /= 10
		fmt.Println(x)
	}
	fmt.Println(revertedNum)
	fmt.Println(x)
	return x == revertedNum || x == revertedNum/10

}

func main() {
	fmt.Println(isPalindrome(1221))
}
