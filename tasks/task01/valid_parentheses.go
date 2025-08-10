package main

import "fmt"

/*
[有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
考察：字符串处理、栈的使用
*/
func isValid(s string) bool {
	stack := []rune{} // 创建一个rune类型的栈，用于存储左括号

	// 定义括号映射关系，用于存储左括号和右括号之间的对应关系
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s { // 遍历字符串中的每个字符
		// 如果是左括号，直接入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else { // 如果是右括号
			// 栈为空或不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != m[char] { // 栈为空
				return false
			}
			// 栈顶元素匹配，出栈
			stack = stack[:len(stack)-1]
		}
	}

	// 最后检查栈是否为空，如果不为空说明还有未匹配的左括号
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
