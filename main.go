package main

import (
	_ "github.com/foolin/gin-template"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"html/template"
	AdminRoutes "main/admin/routes"
	"main/auth/middleware"
	"main/models"
	MainRoutes "main/routes"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "templates",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	// Подключение к базе данных
	models.ConnectDB()

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})
	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeJWT())
	{

		// Маршруты для Admin
		AdminRoutes.Urls(router)
		// Маршруты для main
		MainRoutes.Urls(router)
	}

	// Статика
	router.Static("/assets", "./assets")
	//route.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./assets/resources/favicon.ico")

	// Запуск сервера
	port := "8081"
	router.Run(":" + port)
}
