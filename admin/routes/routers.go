package routes

import (
	"github.com/GoAdminGroup/go-admin/engine"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"main/admin/pages"
)

func Urls(router *gin.Engine, engine *engine.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	//adminRouter := router.Group("/admin")
	//adminRouter.GET("/login", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "login.html", gin.H{
	//		"title": "Авторизация",
	//	})
	//})
	engine.HTML("GET", "/", pages.GetDashBoard)
	engine.HTMLFile("GET", "/admin/hello", ".admin/html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
	//adminRouter.Use(authMiddleware.MiddlewareFunc())
	//{
	//	adminRouter.HTML("GET", "/", pages.GetDashBoard)
	//	adminRouter.HTMLFile("GET", "/admin/hello", ".admin/html/hello.tmpl", map[string]interface{}{
	//		"msg": "Hello world",
	//	})
	//}
}
