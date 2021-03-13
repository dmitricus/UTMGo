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

	"main/admin/models"
	"main/admin/pages"
	"main/admin/tables"
)

func main() {
	startServer()
}

func startServer() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://5310fd7683b54198a2b769f58cbf8042@o465522.ingest.sentry.io/5478277",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{}))

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", ".admin/html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	models.Init(eng.PostgresqlConnection())

	_ = r.Run(":80")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.PostgresqlConnection().Close()
}
