/*
Channel
1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
  - 考察点 ：通道的基本使用、协程间通信。
*/
package main

import (
	"fmt"
	// "time"
	"sync"
)

// producer 生成从1到10的整数并发送到通道
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者发送: %d\n", i)
		ch <- i // 发送数据到通道
	}
	close(ch) // 关闭通道，表示生产结束
}

// consumer 从通道接收整数并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch { // 从通道接收数据，直到通道关闭
		fmt.Printf("消费者接收: %d\n", num)
	}
}

func main() {
	fmt.Println("通道示例")

	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个无缓冲的 int 类型通道
	ch := make(chan int)

	// 启动生产者和消费者协程
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("程序结束")
}
