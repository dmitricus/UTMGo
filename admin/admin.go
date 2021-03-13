package admin

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/GoAdminGroup/go-admin/modules/language"
	_ "github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
	//"github.com/GoAdminGroup/themes/adminlte"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/go-admin/template/types"
)

func InitializeAdmin(context *gin.Context) {
	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "root",
				Name:       "utmgo",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     "postgres",
			},
		},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath:   "/logs/info.log",
		AccessLogPath: "/logs/access.log",
		ErrorLogPath:  "/logs/error.log",
		//ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		Use(context)

	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)
}
