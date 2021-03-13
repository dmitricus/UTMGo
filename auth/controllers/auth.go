package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"main/auth/models"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func AuthenticatorHandler(context *gin.Context) (interface{}, error) {
	var login Login
	var user models.User
	if err := context.ShouldBind(&login); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	//username := login.Username
	//password := login.Password

	//if err := MainModels.DB.Where("Username = ? AND Password = ?", username, password).First(&user).Error; err != nil {
	//	return nil, jwt.ErrFailedAuthentication
	//}

	return &user, nil
}
