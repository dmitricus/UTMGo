package main

import (
	"github.com/gin-gonic/gin"
	AdminRoutes "main/admin/routes"
	"main/auth/middleware"
	AuthRoutes "main/auth/routers"
	"main/models"
	MainRoutes "main/routes"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	//router.GET("/hello", func(c *gin.Context) {
	//	session := sessions.Default(c)
	//
	//	if session.Get("hello") != "world" {
	//		session.Set("hello", "world")
	//		session.Save()
	//	}
	//
	//	c.JSON(200, gin.H{"hello": session.Get("hello")})
	//})
	// Подключение к базе данных
	models.ConnectDB()
	authMiddleware := middleware.AuthMiddleware()
	// Маршруты для Auth
	AuthRoutes.Urls(router, authMiddleware)
	// Маршруты для main
	MainRoutes.Urls(router, authMiddleware)
	// Маршруты для Admin
	AdminRoutes.Urls(router, authMiddleware)

	// Статика
	router.Static("/assets", "./assets")
	//route.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")

	// Запуск сервера
	port := "8081"
	router.Run(":" + port)
}
