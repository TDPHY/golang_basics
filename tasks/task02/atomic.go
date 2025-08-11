/*
锁机制
2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
  - 考察点 ：原子操作、并发数据安全。
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 全局共享计数器，使用int64类型以便原子操作
var atomicCounter int64

// incrementAtomicCounter 使用原子操作对计数器进行递增
func incrementAtomicCounter(wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时通知WaitGroup

	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&atomicCounter, 1) // 原子性地将1加到atomicCounter上
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		go incrementAtomicCounter(&wg)
	}

	// 等待所有协程完成
	wg.Wait()

	duration := time.Since(startTime)

	fmt.Printf("最终计数器值: %d\n", atomicCounter)
	fmt.Printf("执行时间: %v\n", duration)
}
