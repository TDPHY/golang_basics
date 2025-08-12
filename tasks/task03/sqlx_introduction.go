/*
## Sqlx入门
### 题目1：使用SQL扩展库进行查询
- 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
  - 要求 ：
  - 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
  - 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

### 题目2：实现类型安全映射
- 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
  - 要求 ：
  - 定义一个 Book 结构体，包含与 books 表对应的字段。
  - 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Employee 结构体映射 employees 表
type Employee struct {
	gorm.Model
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// Book 结构体映射 books 表
type Book struct {
	gorm.Model
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 题目1：使用SQL扩展库进行查询
func sqlxQueries() {
	// 连接数据库
	db := SqlxConnectDB()

	// employee := Employee{
	// 	Name:       "王五",
	// 	Department: "财务部",
	// 	Salary:     6000.0,
	// }
	// // 执行INSERT语句，使用NamedExec绑定结构体字段
	// if _, err := db.NamedExec(`INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)`, &employee); err != nil {
	// 	log.Fatal("插入记录失败: ", err)
	// }

	// 1. 查询部门为"技术部"的员工信息
	var techEmployees []Employee
	err1 := db.Select(&techEmployees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err1 != nil {
		fmt.Println("查询技术部员工失败:", err1)
	} else {
		fmt.Println("技术部员工:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n",
				emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	}

	// 2. 查询工资最高的员工信息
	var topEmployee Employee
	err2 := db.Get(&topEmployee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err2 != nil {
		fmt.Println("查询最高工资员工失败:", err2)
	} else {
		fmt.Printf("工资最高的员工: ID: %d, Name: %s, Department: %s, Salary: %d\n",
			topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)
	}
}

// 题目2：实现类型安全映射
func typeSafeMapping() {
	db := SqlxConnectDB()

	// book := Book{
	// 	Title:  "Go语言基础",
	// 	Author: "Go",
	// 	Price:  59.9,
	// }
	// // 执行INSERT语句，使用NamedExec绑定结构体字段
	// if _, err := db.NamedExec(`INSERT INTO books (title, author, price) VALUES (:title, :author, :price)`, &book); err != nil {
	// 	log.Fatal("插入记录失败: ", err)
	// }

	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	var expensiveBooks []Book
	err := db.Select(&expensiveBooks, "SELECT id, title, author, price FROM books WHERE price > ?", 50.0)
	if err != nil {
		fmt.Println("查询高价书籍失败:", err)
	} else {
		fmt.Println("价格大于50元的书籍:")
		for _, book := range expensiveBooks {
			fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n",
				book.ID, book.Title, book.Author, book.Price)
		}
	}
}

func SqlxConnectDB() *sqlx.DB {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return nil
	}
	return db
}

func main() {
	fmt.Println("Sqlx入门练习")

	// sqlxQueries()
	typeSafeMapping()
}
