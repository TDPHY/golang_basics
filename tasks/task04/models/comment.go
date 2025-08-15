package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint   `json:"user_id"`
	User    User
	PostID  uint `json:"post_id"`
	Post    Post
}
