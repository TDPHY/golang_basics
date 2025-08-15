package routes

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 用户认证路由
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// 用户路由 (需要认证)
	user := router.Group("/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/:id", controllers.GetUser)
	}

	// 文章路由
	posts := router.Group("/posts")
	posts.Use(middleware.AuthMiddleware())
	{
		posts.GET("/", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		posts.POST("/", controllers.CreatePost)
		posts.PUT("/:id", controllers.UpdatePost)
		posts.DELETE("/:id", controllers.DeletePost)
	}

	// 评论路由
	comments := router.Group("/comments")
	comments.Use(middleware.AuthMiddleware())
	{
		comments.POST("/", controllers.CreateComment)
		comments.GET("/post/:postId", controllers.GetCommentsByPost)
	}
}