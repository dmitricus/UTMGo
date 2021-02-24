package routes

import (
	"github.com/gin-gonic/gin"
	"main/auth/controller"
	"main/auth/service"
	"main/controllers"
	"net/http"
)

func Urls(router *gin.Engine) {
	var loginService = service.StaticLoginService()
	var jwtService = service.JWTAuthService()
	var loginController = controller.LoginHandler(loginService, jwtService)

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	router.GET("/tracks", controllers.GetAllTracks)
	router.POST("/tracks", controllers.CreateTrack)
	router.GET("/tracks/:id", controllers.GetTrack)
	router.PATCH("/tracks/:id", controllers.UpdateTrack)
	router.DELETE("/tracks/:id", controllers.DeleteTrack)
}
