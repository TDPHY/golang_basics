/*
锁机制
1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
  - 考察点 ： sync.Mutex 的使用、并发数据安全。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局共享计数器
var counter int

// 互斥锁
var mutex sync.Mutex

// incrementCounter 对计数器进行递增操作
func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mutex.Lock()   // 加锁
		counter++      // 修改共享数据
		mutex.Unlock() // 解锁
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		go incrementCounter(&wg)
	}

	// 等待所有协程完成
	wg.Wait()

	duration := time.Since(startTime)

	fmt.Printf("最终计数器值: %d\n", counter)
	fmt.Printf("执行时间: %v\n", duration)
}
