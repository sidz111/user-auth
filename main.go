package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/user-auth/config"
	"github.com/sidz111/user-auth/controller"
	"github.com/sidz111/user-auth/model"
	"github.com/sidz111/user-auth/repository"
	"github.com/sidz111/user-auth/routes"
	"github.com/sidz111/user-auth/service"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed")
	} else {
		fmt.Println("Connected ")
	}
	r := gin.Default()
	// config.DB.Migrator(&model.User{})
	config.DB.AutoMigrate(&model.User{})
	userRepo := repository.NewUserRepository(config.DB)
	userServ := service.NewUserService(userRepo)
	userController := controller.NewUserController(userServ)
	authController := controller.NewAuthUserController(userServ)
	route := routes.SetRoutes(userController, authController, r)
	route.Run(":8080")

}
