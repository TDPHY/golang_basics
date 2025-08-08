# Go语言基础教程

这个项目包含了一系列Go语言示例文件。

## 学习顺序

1. [01_hello_world.go](01_hello_world.go) - Go语言基础和Hello World程序
   - Go语言简介
   - 基本程序结构
   - 包和导入
   - main函数

2. [02_variables_and_types.go](02_variables_and_types.go) - 变量和数据类型
   - 变量声明方式
   - 基本数据类型 (int, float, bool, string)
   - 常量和枚举(iota)
   - 复合类型简介 (数组、切片、map)

3. [03_operators_and_control_structures.go](03_operators_and_control_structures.go) - 运算符和控制结构
   - 算术、比较和逻辑运算符
   - 条件语句 (if/else)
   - 分支语句 (switch)
   - 循环语句 (for)

4. [04_functions.go](04_functions.go) - 函数
   - 函数定义和调用
   - 多返回值
   - 可变参数
   - 递归函数
   - 闭包

5. [05_arrays_slices_maps.go](05_arrays_slices_maps.go) - 数组、切片和映射
   - 数组的使用
   - 切片操作 (创建、追加、复制)
   - 映射的使用 (创建、访问、删除)

6. [06_structs_and_methods.go](06_structs_and_methods.go) - 结构体和方法
   - 结构体定义和使用
   - 嵌套结构体
   - 方法定义
   - 值接收者 vs 指针接收者

7. [07_interfaces.go](07_interfaces.go) - 接口
   - 接口定义
   - 接口实现
   - 空接口
   - 类型断言和类型开关

8. [08_error_handling.go](08_error_handling.go) - 错误处理
   - Go中的错误处理理念
   - error接口
   - 自定义错误类型
   - 错误包装和检查

9. [09_goroutines_channels.go](09_goroutines_channels.go) - 并发编程
   - Goroutines基础
   - Channels通信
   - 缓冲通道
   - Select语句

10. [10_packages_modules.go](10_packages_modules.go) - 包和模块管理
    - 包的组织和可见性
    - 模块系统基础
    - 依赖管理

## 如何运行示例

要运行这些示例文件，先安装Go语言环境。然后可以使用以下命令运行单个文件：

```bash
go run 01_hello_world.go
go run 02_variables_and_types.go
# ... 其他文件
```

或者编译成可执行文件：

```bash
go build 01_hello_world.go
./01_hello_world.go  # Linux/Mac
01_hello_world.exe   # Windows
```

## 更多资源

- [Go官方文档](https://golang.org/doc/)
- [Go语言之旅](https://tour.golang.org/)
- [Go标准库文档](https://golang.org/pkg/)
