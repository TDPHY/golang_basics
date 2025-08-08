// Go语言基础教程 - 9. 并发编程：Goroutines和Channels
//
// Go语言内置支持并发编程，通过goroutines和channels实现

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Goroutines
// Goroutines是轻量级线程，由Go运行时管理

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Hello %s (%d)\n", name, i+1)
	}
}

// 2. Channels
// Channels用于goroutines之间的通信和同步

func ping(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("ping %d", i+1)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	close(ch) // 关闭通道
}

func pong(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("pong %d", i+1)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	close(ch)
}

// 3. 带缓冲的通道
func producer(ch chan int, name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("生产者%s生产: %d\n", name, i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(ch chan int, name string) {
	for value := range ch {
		fmt.Printf("消费者%s消费: %d\n", name, value)
		time.Sleep(200 * time.Millisecond)
	}
}

// 4. 使用WaitGroup等待goroutines完成
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时通知WaitGroup
	fmt.Printf("工人%d开始工作\n", id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("工人%d完成工作\n", id)
}

// 5. 选择语句 (select)
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("退出")
			return
		}
	}
}

func main() {
	fmt.Println("=== Goroutines基础 ===")
	
	// 启动goroutine
	go sayHello("Goroutine-1")
	
	// 主goroutine也执行
	sayHello("Main")
	
	// 给goroutine一些时间完成执行
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== Channels基础 ===")
	
	// 创建无缓冲通道
	ch := make(chan string)
	
	// 启动发送者goroutine
	go ping(ch)
	
	// 接收消息
	for msg := range ch {
		fmt.Println("接收到:", msg)
	}
	
	fmt.Println("通道已关闭")
	
	fmt.Println("\n=== 带缓冲的Channels ===")
	
	// 创建带缓冲的通道（缓冲区大小为2）
	bufferedCh := make(chan int, 2)
	
	go producer(bufferedCh, "P1")
	go consumer(bufferedCh, "C1")
	
	// 给goroutines一些时间运行
	time.Sleep(2 * time.Second)
	
	// 关闭通道
	close(bufferedCh)
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== WaitGroup ===")
	
	// 创建WaitGroup
	var wg sync.WaitGroup
	
	// 启动多个worker
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数
		go worker(i, &wg)
	}
	
	// 等待所有worker完成
	wg.Wait()
	fmt.Println("所有工人都已完成工作")
	
	fmt.Println("\n=== Select语句 ===")
	
	c := make(chan int)
	quit := make(chan int)
	
	// 启动fibonacci goroutine
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	
	fibonacci(c, quit)
	
	fmt.Println("\n=== 单向通道 ===")
	
	// 创建通道
	ch1 := make(chan int, 3)
	
	// 发送函数只发送不接收
	go sender(ch1)
	
	// 接收函数只接收不发送
	go receiver(ch1)
	
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== 通道方向示例完成 ===")
}

// 发送者函数，参数是只能发送的通道
func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i * 10
		fmt.Printf("发送: %d\n", i*10)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

// 接收者函数，参数是只能接收的通道
func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Printf("接收: %d\n", value)
	}
}