/*
## SQL语句练习
### 题目1：基本CRUD操作
- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
  - 要求 ：
  - 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
  - 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
  - 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
  - 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

### 题目2：事务语句
- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
  - 要求 ：
  - 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Student
type Student struct {
	gorm.Model
	Name  string `db:"name"`  // 姓名
	Age   int    `db:"age"`   // 年龄
	Grade string `db:"grade"` // 年级
}

// Account
type Account struct {
	gorm.Model
	Balance float64 `db:"balance"` // 账户余额
}

// Transaction（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
type Transaction struct {
	gorm.Model
	FromAccountID uint    `db:"from_account_id"`
	ToAccountID   uint    `db:"to_account_id"`
	Amount        float64 `db:"amount"`
}

// 题目1：基本CRUD操作
func basicCRUD() {
	db := GormInitDB(&Student{})

	// - 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	// db.Create(&Student{Name: "李四", Age: 14, Grade: "二年级"})

	// - 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Student
	db.Find(&students, "age > ?", 18)
	fmt.Printf("年龄大于 18 的学生信息: %v\n", students)
	// - 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Student{}).Where("name LIKE ?", "张三%").Update("grade", "四年级")
	// - 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Delete(&Student{}, "age < ?", 15)

}

// 题目2：事务语句
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务。
func transactionExample() {
	db := GormInitDB(&Account{}, &Transaction{})

	// // 创建两个账户
	// accountA := Account{Balance: 500.0}
	// accountB := Account{Balance: 500.0}
	// // 插入账户
	// db.Create(&accountA)
	// db.Create(&accountB)
	// // 查询账户
	// db.Find(&accountA)

	// 事务开始
	err := db.Transaction(func(tx *gorm.DB) error {
		var accountA, accountB Account
		if err := tx.Where("id = ?", 1).First(&accountA).Error; err != nil {
			return err
		}
		fmt.Println("账户A余额:", accountA.Balance)
		if err := tx.Where("id = ?", 2).First(&accountB).Error; err != nil {
			return err
		}
		fmt.Println("账户B余额:", accountB.Balance)

		// 检查账户 A 的余额是否足够，如果余额不足，则回滚事务
		if accountA.Balance < 100.0 {
			return errors.New("账户 A 的余额不足")
		}

		// 从账户 A 扣除 100 元
		accountA.Balance -= 100.0
		if err := tx.Save(&accountA).Error; err != nil {
			return err // 出错自动回滚
		}
		fmt.Println("账户A扣除 100 元")

		// 向账户 B 添加 100 元
		accountB.Balance += 100.0
		if err := tx.Save(&accountB).Error; err != nil {
			return err // 出错自动回滚
		}
		fmt.Println("账户B添加 100 元")

		// 在 transactions 表中记录该笔转账信息
		transaction := Transaction{
			Amount:        100.0,
			FromAccountID: accountA.ID,
			ToAccountID:   accountB.ID,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err // 出错自动回滚
		}
		fmt.Println("记录该笔转账信息")

		// 所有操作都成功了，事务提交
		return nil

	})

	if err != nil {
		log.Fatal("事务执行失败，已回滚：", err)
	} else {
		log.Println("事务执行成功")
	}
}

func GormConnectDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return nil
	}
	return db
}

func GormInitDB(entities ...interface{}) *gorm.DB {
	db := GormConnectDB()
	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Printf("Failed to migrate database: %v", err)
			panic("Failed to migrate database: " + err.Error())
		}
	}
	return db
}

func main() {
	fmt.Println("SQL语句练习")

	// basicCRUD()

	transactionExample()
}
