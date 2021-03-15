package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetEmailServersTable return the model of table email server.
func GetEmailServersTable(ctx *context.Context) (emailServersTable table.Table) {

	emailServersTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := emailServersTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Хост", "email_host", db.Varchar)
	info.AddField("Порт", "email_port", db.Int)
	info.AddField("SSL", "email_use_ssl", db.Boolean).FieldBool("true", "false")
	info.AddField("TLS", "email_use_tls", db.Boolean).FieldBool("true", "false")
	info.AddField("Способ отправки писем", "sending_method", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			if model.Value == "smtp" {
				return "SMTP сервер"
			}
			if model.Value == "unisender_api" {
				return "UniSender API"
			}
			return "unknown"
		})
	info.AddField("Основной сервер", "main_server", db.Boolean).FieldBool("true", "false")
	info.AddField("Сервер активен", "is_active", db.Boolean).FieldBool("true", "false")

	info.SetTable("email_servers").SetTitle("Почтовые сервера").SetDescription("Почтовые сервера")

	formList := emailServersTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("FROM", "email_default_from", db.Varchar, form.Text)
	formList.AddField("Хост", "email_host", db.Varchar, form.Text)
	formList.AddField("порт", "email_port", db.Int, form.Text)
	formList.AddField("Имя пользователя", "email_username", db.Varchar, form.Text)
	formList.AddField("Пароль", "email_password", db.Varchar, form.Password)
	formList.AddField("SSL", "email_use_ssl", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "true"},
			{Value: "false"},
		}).FieldDefault("false")
	formList.AddField("TLS", "email_use_tls", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "true"},
			{Value: "false"},
		}).FieldDefault("false")
	formList.AddField("Тихое подавление ошибок", "email_fail_silently", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "true"},
			{Value: "false"},
		}).FieldDefault("false")
	formList.AddField("Тайм-аут", "email_timeout", db.Int, form.Number).FieldDefault("0")
	formList.AddField("Сертификат", "email_ssl_cert_file", db.Varchar, form.File)
	formList.AddField("Файл ключа", "email_ssl_key_file", db.Varchar, form.File)
	formList.AddField("API KEY", "api_key", db.Varchar, form.Text)
	formList.AddField("Имя пользователя для авторизации в API", "api_username", db.Varchar, form.Text)
	formList.AddField("Email адрес для отправки через API", "api_from_email", db.Varchar, form.Text)
	formList.AddField("Имя перед адресом для отправки через API", "api_from_name", db.Varchar, form.Text)
	formList.AddField("Способ отправки писем", "sending_method", db.Varchar, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "SMTP сервер", Value: "smtp"},
			{Text: "UniSender API", Value: "unisender_api"},
		})
	formList.AddField("Основной сервер", "main_server", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "true"},
			{Value: "false"},
		}).FieldDefault("false")
	formList.AddField("Сервер активен", "is_active", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "true"},
			{Value: "false"},
		}).FieldDefault("false")
	formList.AddField("Предпочтительней для доменов", "preferred_domains", db.Tinyint, form.SelectSingle)

	formList.SetTable("email_servers").SetTitle("Почтовый сервер").SetDescription("Почтовый сервер")

	return
}
