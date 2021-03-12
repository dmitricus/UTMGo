package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
	"time"
)

func GetVersion(context *gin.Context) {
	var now time.Time
	models.DB.Raw("SELECT NOW()").Scan(&now)
	context.JSON(http.StatusOK, gin.H{
		"version":         "",
		"build_time":      "",
		"commit":          "",
		"db_datetime":     now,
		"server_datetime": time.Now(),
	})
}
