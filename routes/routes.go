package routes

import (
	"story-plateform/controllers"
	"story-plateform/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = false

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "story-plateform backend working",
		})
	})

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.AuthMiddleware())

	// auth routes
	router.POST("/login", controllers.Login)
	protectedRoutes.POST("/logout", controllers.Logout)

	// users routes
	protectedRoutes.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)

	// stories routes
	protectedRoutes.GET("/stories", controllers.GetStories)
	protectedRoutes.GET("/stories/:id", controllers.GetStoryById)
	protectedRoutes.POST("/stories", controllers.CreateStory)
	protectedRoutes.PUT("/stories/:id", controllers.UpdateStory)
	protectedRoutes.DELETE("/stories/:id", controllers.DeleteStory)

	// websocket route for collabaration
	protectedRoutes.GET("/ws", controllers.WebsocketHanler)

	return router
}
