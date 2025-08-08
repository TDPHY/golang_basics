// Go语言基础教程 - 8. 错误处理
//
// Go语言通过返回错误值的方式处理错误，而不是使用异常

package main

import (
	"errors"
	"fmt"
	"math"
)

// 1. Go中的错误类型
// error是一个内置接口，定义如下:
// type error interface {
//     Error() string
// }

// 2. 创建自定义错误
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// 使用errors.New创建简单错误
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

// 3. 使用fmt.Errorf创建格式化错误
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %f by zero", a)
	}
	return a / b, nil
}

// 4. 自定义错误类型
type MyError struct {
	Code    int
	Message string
}

// 实现error接口
func (e *MyError) Error() string {
	return fmt.Sprintf("错误代码: %d, 错误信息: %s", e.Code, e.Message)
}

func DoSomething(success bool) error {
	if !success {
		return &MyError{
			Code:    1001,
			Message: "操作失败",
		}
	}
	return nil
}

// 5. 错误包装 (Go 1.13+)
func ProcessData(data string) error {
	// 模拟处理过程中的错误
	if data == "" {
		return fmt.Errorf("处理数据时出错: %w", errors.New("数据为空"))
	}
	return nil
}

// 6. 错误检查函数
func IsEmptyDataError(err error) bool {
	// 注意：实际应用中应使用errors.Is等函数进行更准确的判断
	return err != nil && err.Error() == "数据为空"
}

func main() {
	fmt.Println("=== 基本错误处理 ===")
	
	// 处理可能返回错误的函数
	result, err := Sqrt(16)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("平方根: %.2f\n", result)
	}
	
	// 处理错误情况
	result, err = Sqrt(-4)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("平方根: %.2f\n", result)
	}
	
	fmt.Println("\n=== 格式化错误 ===")
	// 除法运算
	quotient, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}
	
	quotient, err = Divide(5, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("5 / 0 = %.2f\n", quotient)
	}
	
	fmt.Println("\n=== 自定义错误类型 ===")
	// 成功情况
	err = DoSomething(true)
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)
	} else {
		fmt.Println("操作成功")
	}
	
	// 失败情况
	err = DoSomething(false)
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)
		// 类型断言获取具体错误信息
		if myErr, ok := err.(*MyError); ok {
			fmt.Printf("错误代码: %d\n", myErr.Code)
			fmt.Printf("错误信息: %s\n", myErr.Message)
		}
	} else {
		fmt.Println("操作成功")
	}
	
	fmt.Println("\n=== 错误包装和检查 ===")
	// 正常情况
	err = ProcessData("有效数据")
	if err != nil {
		fmt.Printf("处理错误: %v\n", err)
	}
	
	// 错误情况
	err = ProcessData("")
	if err != nil {
		fmt.Printf("处理错误: %v\n", err)
		// 检查特定错误
		if IsEmptyDataError(err) {
			fmt.Println("检测到数据为空错误")
		}
	}
	
	fmt.Println("\n=== 多重错误检查 ===")
	// 模拟文件操作中的多重错误处理
	files := []string{"file1.txt", "", "file2.txt", "nonexistent.txt"}
	
	for _, file := range files {
		fmt.Printf("\n处理文件: '%s'\n", file)
		if file == "" {
			fmt.Println("错误: 文件名为空")
			continue
		}
		if file == "nonexistent.txt" {
			fmt.Println("错误: 文件不存在")
			continue
		}
		fmt.Println("文件处理成功")
	}
	
	fmt.Println("\n=== 错误处理最佳实践 ===")
	// 1. 总是检查错误
	// 2. 尽早处理错误
	// 3. 提供有意义的错误信息
	// 4. 不要忽略错误
	
	// 示例：模拟数据库连接
	dbConnections := map[string]string{
		"mysql":    "mysql://localhost:3306",
		"postgres": "postgres://localhost:5432",
		"invalid":  "",
	}
	
	for dbType, connStr := range dbConnections {
		fmt.Printf("\n连接数据库: %s\n", dbType)
		if connStr == "" {
			fmt.Printf("错误: 无效的%s连接字符串\n", dbType)
			continue
		}
		fmt.Printf("成功连接到%s数据库: %s\n", dbType, connStr)
	}
}