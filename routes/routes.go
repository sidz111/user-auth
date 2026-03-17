package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/user-auth/controller"
	"github.com/sidz111/user-auth/middleware"
)

func SetRoutes(usersController *controller.UserController, authController *controller.AuthController, router *gin.Engine) *gin.Engine {
	user := router.Group("users")
	router.POST("/login", authController.Login)
	{
		user.POST("/", usersController.CreateUser)
		user.GET("/:id", middleware.AuthMiddleware(), usersController.GetUser)
		user.GET("/", middleware.AuthMiddleware(), usersController.GetAllUsers)
		user.PUT("/", middleware.AuthMiddleware(), usersController.UpdateUser)
		user.DELETE("/:id", middleware.AuthMiddleware(), usersController.DeleteUser)
	}
	return router
}
