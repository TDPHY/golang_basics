/*
Goroutine
2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
  - 考察点 ：协程原理、并发任务调度。
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Task 代表一个任务
type Task func()

// simulateTask 模拟一个耗时任务
func simulateTask(name string) Task {
	return func() {
		// 模拟随机执行时间 (100ms - 1000ms)
		duration := time.Duration(rand.Intn(900)+100) * time.Millisecond
		fmt.Printf("任务 %s 开始执行\n", name)
		time.Sleep(duration)
		fmt.Printf("任务 %s 执行完成，耗时: %v\n", name, duration)
	}
}

// executeTasks 并发执行任务并统计执行时间
func executeTasks(tasks []Task) {
	startTime := time.Now()

	// 创建一个通道用于等待所有任务完成
	done := make(chan bool, len(tasks))

	// 并发执行所有任务
	for i, task := range tasks {
		go func(index int, t Task) {
			defer func() { done <- true }() // 任务完成后发送信号
			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)
			fmt.Printf("任务 %d 实际执行时间: %v\n", index, duration)
		}(i, task)
	}

	// 等待所有任务完成
	for i := 0; i < len(tasks); i++ {
		<-done
	}

	totalDuration := time.Since(startTime)
	fmt.Printf("所有任务执行完成，总耗时: %v\n", totalDuration)
}

func main() {
	fmt.Println("=== 任务调度器开始 ===")

	// 创建一组任务
	tasks := []Task{
		simulateTask("下载文件"),
		simulateTask("处理数据"),
		simulateTask("发送邮件"),
		simulateTask("生成报告"),
	}

	// 执行任务
	executeTasks(tasks)
}

// package main

// import (
//     "fmt"
//     "math/rand"
//     "sync"
//     "time"
// )

// /*
// Goroutine
// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//   - 考察点 ：协程原理、并发任务调度。
// */

// // Task 任务结构体
// type Task struct {
//     Name     string
//     Function func() error
// }

// // TaskResult 任务执行结果
// type TaskResult struct {
//     TaskName   string
//     Duration   time.Duration
//     Error      error
// }

// // TaskScheduler 任务调度器
// type TaskScheduler struct {
//     tasks   []Task
//     results chan TaskResult
//     wg      sync.WaitGroup
// }

// // NewTaskScheduler 创建新的任务调度器
// func NewTaskScheduler() *TaskScheduler {
//     return &TaskScheduler{
//         tasks:   make([]Task, 0),
//         results: make(chan TaskResult, 100), // 缓冲通道
//     }
// }

// // AddTask 添加任务
// func (ts *TaskScheduler) AddTask(name string, fn func() error) {
//     ts.tasks = append(ts.tasks, Task{
//         Name:     name,
//         Function: fn,
//     })
// }

// // executeTask 执行单个任务并统计时间
// func (ts *TaskScheduler) executeTask(task Task) {
//     defer ts.wg.Done()

//     start := time.Now()
//     err := task.Function()
//     duration := time.Since(start)

//     // 发送结果到通道
//     ts.results <- TaskResult{
//         TaskName: task.Name,
//         Duration: duration,
//         Error:    err,
//     }
// }

// // Run 并发执行所有任务
// func (ts *TaskScheduler) Run() []TaskResult {
//     // 启动所有任务
//     for _, task := range ts.tasks {
//         ts.wg.Add(1)
//         go ts.executeTask(task)
//     }

//     // 启动结果收集协程
//     go func() {
//         ts.wg.Wait()
//         close(ts.results)
//     }()

//     // 收集所有结果
//     var results []TaskResult
//     for result := range ts.results {
//         results = append(results, result)
//     }

//     return results
// }

// // 示例任务函数
// func task1() error {
//     fmt.Printf("任务1开始执行...\n")
//     time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机延迟
//     fmt.Printf("任务1执行完成\n")
//     return nil
// }

// func task2() error {
//     fmt.Printf("任务2开始执行...\n")
//     time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机延迟
//     fmt.Printf("任务2执行完成\n")
//     return nil
// }

// func task3() error {
//     fmt.Printf("任务3开始执行...\n")
//     time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机延迟
//     fmt.Printf("任务3执行完成\n")
//     return nil
// }

// func task4() error {
//     fmt.Printf("任务4开始执行...\n")
//     time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机延迟
//     fmt.Printf("任务4执行完成\n")
//     return nil
// }

// func taskWithError() error {
//     fmt.Printf("任务(带错误)开始执行...\n")
//     time.Sleep(500 * time.Millisecond)
//     fmt.Printf("任务(带错误)执行完成，但返回错误\n")
//     return fmt.Errorf("这是一个示例错误")
// }

// func main() {
//     fmt.Println("=== 任务调度器演示 ===")

//     // 创建任务调度器
//     scheduler := NewTaskScheduler()

//     // 添加任务
//     scheduler.AddTask("数据处理", task1)
//     scheduler.AddTask("网络请求", task2)
//     scheduler.AddTask("文件读写", task3)
//     scheduler.AddTask("数据库操作", task4)
//     scheduler.AddTask("错误处理示例", taskWithError)

//     // 记录总执行时间
//     startTime := time.Now()

//     // 并发执行所有任务
//     results := scheduler.Run()

//     totalTime := time.Since(startTime)

//     // 输出结果统计
//     fmt.Println("\n=== 任务执行结果 ===")
//     for _, result := range results {
//         if result.Error != nil {
//             fmt.Printf("任务: %s | 执行时间: %v | 错误: %v\n",
//                 result.TaskName, result.Duration, result.Error)
//         } else {
//             fmt.Printf("任务: %s | 执行时间: %v\n",
//                 result.TaskName, result.Duration)
//         }
//     }

//     fmt.Printf("\n总执行时间: %v\n", totalTime)
//     fmt.Println("注意：总执行时间远小于各任务执行时间之和，说明任务是并发执行的")
// }
