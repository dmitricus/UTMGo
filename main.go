package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                      // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"main/auth/middleware"

	"main/admin/models"

	AdminRoutes "main/admin/routes"
	"main/admin/tables"
	ApiRoutes "main/api/routers"
	AuthRoutes "main/auth/routers"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	eng := engine.Default()

	// Настройки Сентри
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://5310fd7683b54198a2b769f58cbf8042@o465522.ingest.sentry.io/5478277",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	r.Use(sentrygin.New(sentrygin.Options{}))

	template.AddComp(chartjs.NewChart())

	// Конфиг
	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}
	r.Static("/uploads", "./uploads")

	authMiddleware := middleware.AuthMiddleware()
	// Маршруты для Auth
	AuthRoutes.Urls(r, authMiddleware)
	// Маршруты для Admin
	AdminRoutes.Urls(r, eng, authMiddleware)
	// Маршруты для Api
	ApiRoutes.Urls(r, authMiddleware)

	models.Init(eng.PostgresqlConnection())

	_ = r.Run(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.PostgresqlConnection().Close()
}
