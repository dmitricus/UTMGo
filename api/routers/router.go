package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"main/api/controllers"
	"net/http"
)

func Urls(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	api := router.Group("/api")
	api.GET("/version", controllers.GetVersion)
	api.GET("/send", func(context *gin.Context) {
		controllers.SendEmail(
			"user1@example.com",
			"Тестовое письмо",
			"тест Тест ТЕСТ",
			"20100",
			"info.msg",
		)
		context.JSON(http.StatusOK, gin.H{"user1@example.com": "OK"})
	})
}
