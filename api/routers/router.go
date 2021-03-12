package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"main/api/controllers"
)

func Urls(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	api := router.Group("/api")
	api.GET("/version", controllers.GetVersion)
}
