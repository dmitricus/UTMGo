package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"main/auth/middleware"
)

func Urls(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", middleware.HelloHandler)
	}

	//router.POST("/login", controllers.LoginHandler)
	//router.GET("/logout", controllers.LogoutHandler)
	//router.POST("/users/", controllers2.CreateUser)
}
