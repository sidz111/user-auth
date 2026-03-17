package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/user-auth/config"
	"github.com/sidz111/user-auth/model"
	"github.com/sidz111/user-auth/service"
	"github.com/sidz111/user-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	serv service.UserService
}

func NewAuthUserController(serv service.UserService) *AuthController {
	return &AuthController{serv: serv}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user model.User
	var foundUser model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := config.DB.Model(&model.User{}).Where("id = ?", user.ID).First(&foundUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return
	}
	tokenString, err := utils.GenerateJWT(foundUser.Username, foundUser.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}
