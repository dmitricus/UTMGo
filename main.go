package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	AdminRoutes "main/admin/routes"
	ApiRoutes "main/api/routers"
	"main/auth/middleware"
	AuthRoutes "main/auth/routers"
	MainControllers "main/controllers"
	"main/models"
	MainRoutes "main/routes"
	"net/http"
)

func main() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://5310fd7683b54198a2b769f58cbf8042@o465522.ingest.sentry.io/5478277",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(sentrygin.New(sentrygin.Options{}))

	// Подключение к базе данных
	models.ConnectDB()
	authMiddleware := middleware.AuthMiddleware()
	// Маршруты для Auth
	AuthRoutes.Urls(router, authMiddleware)
	// Маршруты для main
	MainRoutes.Urls(router, authMiddleware)
	// Маршруты для Admin
	AdminRoutes.Urls(router, authMiddleware)
	// Маршруты для Api
	ApiRoutes.Urls(router, authMiddleware)

	// Статика
	router.Static("/assets", "./assets")
	//route.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")

	router.GET("/send", func(context *gin.Context) {
		MainControllers.SendEmail("user1@example.com", "20100", "password.msg")
		context.JSON(http.StatusOK, gin.H{"user1@example.com": "OK"})
	})

	// Запуск сервера
	port := "8081"
	router.Run(":" + port)
}
