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
		controllers.SendEmail("borodulin_da@srvhub.ru", "20100", "password.msg")
		context.JSON(http.StatusOK, gin.H{"user1@example.com": "OK"})
	})
}
