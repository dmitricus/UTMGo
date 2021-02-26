package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
)

func Urls(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tracks", controllers.GetAllTracks)
	router.POST("/tracks", controllers.CreateTrack)
	router.GET("/tracks/:id", controllers.GetTrack)
	router.PATCH("/tracks/:id", controllers.UpdateTrack)
	router.DELETE("/tracks/:id", controllers.DeleteTrack)
}
