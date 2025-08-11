/*
Channel
2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
  - 考察点 ：通道的缓冲机制。
*/
package main

import (
	"fmt"
	// "time"
	"sync"
)

// producer 向缓冲通道发送100个整数
func producer2(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i // 发送数据到缓冲通道
		fmt.Printf("生产者发送: %d\n", i)
	}

	// 关闭通道
	close(ch)
}

// consumer 从缓冲通道接收整数并打印
func consumer2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("消费者接收: %d\n", num)
	}
}

func main() {
	fmt.Println("通道示例")

	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个缓冲通道，缓冲区大小为10
	ch := make(chan int, 10)

	// 启动生产者和消费者协程
	go producer2(ch, &wg)
	go consumer2(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	// 判断通道是否关闭
	_, isClosed := <-ch
	if isClosed == false {
		fmt.Println("通道已关闭")
	} else {
		fmt.Println("通道未关闭")
	}
	fmt.Println("程序结束")
}
