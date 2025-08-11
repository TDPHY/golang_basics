/*
Goroutine
1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
  - 考察点 ： go 关键字的使用、协程的并发执行。
*/
package main

import (
	"fmt"
	// "time"
	"sync"
)

// // 打印奇数的协程函数1
// func printOdd() {
// 	fmt.Println("奇数协程开始执行...")
// 	for i := 1; i <= 10; i += 2 {
// 		fmt.Printf("奇数: %d\n", i)
// 		time.Sleep(100 * time.Millisecond) // 添加短暂延迟以便观察输出
// 	}
// 	fmt.Println("奇数协程执行完毕")
// }

// // 打印偶数的协程函数1
// func printEven() {
// 	fmt.Println("偶数协程开始执行...")
// 	for i := 2; i <= 10; i += 2 {
// 		fmt.Printf("偶数: %d\n", i)
// 		time.Sleep(100 * time.Millisecond) // 添加短暂延迟以便观察输出
// 	}
// 	fmt.Println("偶数协程执行完毕")
// }

// 打印奇数的协程函数2
func printOddNumbers(wg *sync.WaitGroup) {
	fmt.Println("奇数协程开始执行...")
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
	fmt.Println("奇数协程执行完毕")

	defer wg.Done() // 协程结束时通知WaitGroup
}

// 打印偶数的协程函数2
func printEvenNumbers(wg *sync.WaitGroup) {
	fmt.Println("偶数协程开始执行...")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
	fmt.Println("偶数协程执行完毕")

	defer wg.Done() // 协程结束时通知WaitGroup
}

func main() {
	fmt.Println("主程序开始执行")

	// // 使用 go 关键字启动两个协程
	// go printOdd()
	// go printEven()
	// // 主协程等待一段时间，确保子协程执行完毕
	// time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	// Add(n)：添加需要等待的协程数量
	wg.Add(2) // 添加两个等待任务

	// 使用go关键字启动两个协程
	go printOddNumbers(&wg)
	go printEvenNumbers(&wg)

	// Wait()：阻塞等待所有协程完成
	wg.Wait()

	fmt.Println("主程序执行完毕")
}
