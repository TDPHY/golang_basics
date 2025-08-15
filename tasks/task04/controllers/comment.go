package controllers

import (
	"blog/config"
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置评论的用户ID
	comment.UserID = userID.(uint)

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetCommentsByPost 获取某篇文章的所有评论
func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("postId")
	// 验证postID是否为有效数字
	_, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var comments []models.Comment

	if err := config.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
