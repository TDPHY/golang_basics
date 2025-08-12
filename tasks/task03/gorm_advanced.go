/*
## 进阶gorm
### 题目1：模型定义
- 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
  - 要求 ：
  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
  - 编写Go代码，使用Gorm创建这些模型对应的数据库表。

### 题目2：关联查询
- 基于上述博客系统的模型定义。
  - 要求 ：
  - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
  - 编写Go代码，使用Gorm查询评论数量最多的文章信息。

### 题目3：钩子函数
- 继续使用博客系统的模型。
  - 要求 ：
  - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
  - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User (用户)
type User struct {
	gorm.Model
	Name      string
	Posts     []Post // 一个用户可以有多篇文章
	PostCount uint   // 文章数量
}

// Post (文章)
type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint      // 外键关联到User
	Comments      []Comment // 一篇文章可以有多个评论
	CommentCount  uint      // 评论数量
	CommentStatus string    // 评论状态
}

// Comment (评论)
type Comment struct {
	gorm.Model
	Content string
	PostID  uint // 外键关联到Post
}

// 题目1：模型定义和关系
func defineModels() *gorm.DB {
	db := GormInitDB(&User{}, &Post{}, &Comment{})

	log.Println("数据库表创建成功！")

	return db
}

// 题目2：关联查询
func associationQueries(db *gorm.DB) {
	// 1. 查询某个用户发布的所有文章及其对应的评论信息
	var user User

	// 查询ID为1的用户
	db.Preload("Posts.Comments").First(&user, 1)

	fmt.Printf("用户 %s 的文章:\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("- 文章: %s\n", post.Title)
		fmt.Println("  评论:")
		for _, comment := range post.Comments {
			fmt.Printf("  - %s\n", comment.Content)
		}
	}

	// 2. 查询评论数量最多的文章信息
	var postWithMostComments Post
	db.Raw(`
		SELECT posts.*, COUNT(comments.id) as comment_count
		FROM posts
		LEFT JOIN comments ON posts.id = comments.post_id
		GROUP BY posts.id
		ORDER BY comment_count DESC
		LIMIT 1
	`).Scan(&postWithMostComments)

	// db.Model(&Post{}).
	// 	Joins("left join comments on comments.post_id = posts.id").
	// 	Group("posts.id").
	// 	Select("posts.*, count(comments.id) as comment_count").
	// 	Order("comment_count desc").
	// 	Limit(1).
	// 	Scan(&postWithMostComments)

	fmt.Printf("评论数量最多的文章: %s\n", postWithMostComments.Title)
}

// 题目3：钩子函数
// 为 Post 模型添加钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 在文章创建时自动更新用户的文章数量统计字段
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// 为 Comment 模型添加钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 在评论删除时检查文章的评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	// 如果评论数量为 0，则更新文章的评论状态为 "无评论"

	if commentCount == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

// 创建示例数据
func createSampleData(db *gorm.DB) {
	// 创建用户
	user := User{Name: "张三"}
	db.Create(&user)

	// 创建文章
	post := Post{Title: "Go语言学习", Content: "Go语言是一门强大的编程语言", UserID: user.ID}
	db.Create(&post)

	// 创建评论
	comment1 := Comment{Content: "很好的文章", PostID: post.ID}
	comment2 := Comment{Content: "学到了很多", PostID: post.ID}
	db.Create(&comment1)
	db.Create(&comment2)

	fmt.Println("示例数据创建完成")
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
	fmt.Println("GORM进阶练习")

	db := GormConnectDB()

	// 定义模型
	// db := defineModels()

	// 创建示例数据
	// createSampleData(db)

	// // 关联查询
	associationQueries(db)
}
