package controllers

import (
	"blog/config"
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser 获取用户信息
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	// 验证userID是否为有效数字
	_, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 不返回密码字段
	user.Password = ""
	c.JSON(http.StatusOK, user)
}