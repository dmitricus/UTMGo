package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func Urls(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
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
	adminRouter := router.Group("/admin")
	adminRouter.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Авторизация",
		})
	})

	adminRouter.Use(authMiddleware.MiddlewareFunc())
	{
		adminRouter.GET("/", func(ctx *gin.Context) {
			//render with master
			ctx.HTML(http.StatusOK, "index", gin.H{
				"title": "Index title!",
				"add": func(a int, b int) int {
					return a + b
				},
			})
		})
		adminRouter.GET("/page", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "page", gin.H{
				"title": "Page file title!!",
			})
		})
	}
}
