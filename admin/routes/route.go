package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Urls(route *gin.Engine) {
	route.GET("/admin", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello Admin!"})
	})
}
