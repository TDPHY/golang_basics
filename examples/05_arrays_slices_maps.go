// Go语言基础教程 - 5. 数组、切片和映射
//
// 数组、切片和映射是Go语言中重要的数据结构

package main

import "fmt"

func main() {
	// 1. 数组 (Array)
	// 数组是固定长度的序列，长度是类型的一部分
	fmt.Println("=== 数组 ===")
	
	// 声明和初始化数组
	var arr1 [5]int                    // 零值初始化
	arr2 := [5]int{1, 2, 3, 4, 5}     // 指定所有元素
	arr3 := [5]int{1, 2}              // 部分初始化，其余为零值
	arr4 := [...]int{1, 2, 3, 4, 5}   // 让编译器计算长度
	
	fmt.Printf("arr1: %v, 长度: %d\n", arr1, len(arr1))
	fmt.Printf("arr2: %v, 长度: %d\n", arr2, len(arr2))
	fmt.Printf("arr3: %v, 长度: %d\n", arr3, len(arr3))
	fmt.Printf("arr4: %v, 长度: %d\n", arr4, len(arr4))
	
	// 访问和修改数组元素
	arr2[0] = 10
	fmt.Printf("修改后的arr2: %v\n", arr2)
	
	// 多维数组
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}
	fmt.Printf("二维数组matrix:\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", matrix[i])
	}
	
	// 2. 切片 (Slice)
	// 切片是动态数组，提供了更强大的数组操作功能
	fmt.Println("\n=== 切片 ===")
	
	// 从数组创建切片
	slice1 := arr2[1:4]  // 从索引1到3（不包括4）
	slice2 := arr2[:3]   // 从开头到索引2
	slice3 := arr2[2:]   // 从索引2到结尾
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("slice1 (arr2[1:4]): %v\n", slice1)
	fmt.Printf("slice2 (arr2[:3]): %v\n", slice2)
	fmt.Printf("slice3 (arr2[2:]): %v\n", slice3)
	
	// 直接创建切片
	slice4 := []int{1, 2, 3, 4, 5}  // 不指定长度，创建切片
	fmt.Printf("slice4: %v, 长度: %d, 容量: %d\n", slice4, len(slice4), cap(slice4))
	
	// 使用make创建切片
	slice5 := make([]int, 5)        // 长度为5的切片，元素为零值
	slice6 := make([]int, 3, 10)    // 长度为3，容量为10的切片
	fmt.Printf("slice5: %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))
	fmt.Printf("slice6: %v, 长度: %d, 容量: %d\n", slice6, len(slice6), cap(slice6))
	
	// 切片操作
	fmt.Println("\n--- 切片操作 ---")
	slice7 := []int{1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", slice7)
	
	// 添加元素
	slice7 = append(slice7, 6)
	fmt.Printf("添加一个元素后: %v\n", slice7)
	
	slice7 = append(slice7, 7, 8, 9)
	fmt.Printf("添加多个元素后: %v\n", slice7)
	
	// 连接两个切片
	slice8 := []int{10, 11, 12}
	slice7 = append(slice7, slice8...)
	fmt.Printf("连接slice8后: %v\n", slice7)
	
	// 复制切片
	slice9 := make([]int, len(slice7))
	copy(slice9, slice7)
	fmt.Printf("复制后的slice9: %v\n", slice9)
	
	// 切片共享底层数组
	fmt.Println("\n--- 切片共享底层数组 ---")
	original := []int{0, 1, 2, 3, 4, 5}
	sliceA := original[1:4]  // [1, 2, 3]
	sliceB := original[2:5]  // [2, 3, 4]
	fmt.Printf("original: %v\n", original)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v\n", sliceB)
	
	// 修改sliceA会影响sliceB和original
	sliceA[0] = 100
	fmt.Printf("修改sliceA[0]=100后:\n")
	fmt.Printf("original: %v\n", original)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v\n", sliceB)
	
	// 3. 映射 (Map)
	// Map是键值对的无序集合
	fmt.Println("\n=== 映射 ===")
	
	// 创建map
	var map1 map[string]int           // 声明但未初始化，值为nil
	map2 := make(map[string]int)      // 使用make初始化
	map3 := map[string]int{           // 字面量初始化
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	
	// 检查map是否初始化
	if map1 == nil {
		fmt.Println("map1未初始化，是nil")
		map1 = make(map[string]int) // 初始化
	}
	
	// 添加和修改元素
	map2["red"] = 1
	map2["green"] = 2
	map2["blue"] = 3
	map2["red"] = 10 // 修改已存在的键
	
	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("map3: %v\n", map3)
	
	// 访问元素
	fmt.Printf("map3中apple的数量: %d\n", map3["apple"])
	fmt.Printf("map3中grape的数量: %d\n", map3["grape"]) // 不存在的键返回零值
	
	// 检查键是否存在
	if value, exists := map3["banana"]; exists {
		fmt.Printf("banana存在，数量为: %d\n", value)
	} else {
		fmt.Println("banana不存在")
	}
	
	// 删除元素
	delete(map3, "orange")
	fmt.Printf("删除orange后map3: %v\n", map3)
	
	// 遍历map
	fmt.Println("\n--- 遍历map ---")
	for key, value := range map3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	
	// 只遍历键
	for key := range map3 {
		fmt.Printf("Key: %s\n", key)
	}
	
	// 只遍历值
	for _, value := range map3 {
		fmt.Printf("Value: %d\n", value)
	}
}