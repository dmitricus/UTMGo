package routes

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func Urls(router *gin.Engine) {
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "admin/templates",
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
	admin := router.Group("/admin")
	admin.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})
	admin.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Авторизация",
		})
	})
	admin.GET("/page", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "page", gin.H{
			"title": "Page file title!!",
		})
	})
}
