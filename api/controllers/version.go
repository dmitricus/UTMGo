package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"main/admin/models"
)

func GetVersion(context *gin.Context) {
	var now time.Time
	models.ORM.Raw("SELECT NOW()").Scan(&now)
	context.JSON(http.StatusOK, gin.H{
		"version":         "",
		"build_time":      "",
		"commit":          "",
		"db_datetime":     now,
		"server_datetime": time.Now(),
	})
}
