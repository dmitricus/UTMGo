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
	//info.AddField("EmailDefaultFrom", "email_default_from", db.Varchar)
	info.AddField("Хост", "email_host", db.Varchar)
	info.AddField("Порт", "email_port", db.Int)
	//info.AddField("EmailUsername", "email_username", db.Varchar)
	info.AddField("SSL", "email_use_ssl", db.Boolean)
	info.AddField("TLS", "email_use_tls", db.Boolean)
	//info.AddField("EmailFailSilently", "email_fail_silently", db.Varchar)
	//info.AddField("EmailTimeout", "email_timeout", db.Int)
	//info.AddField("EmailSSLCertFile", "email_ssl_cert_file", db.Varchar)
	//info.AddField("EmailSSLKeyfile", "email_ssl_key_file", db.Varchar)
	//info.AddField("ApiKey", "api_key", db.Varchar)
	//info.AddField("ApiUsername", "api_username", db.Varchar)
	//info.AddField("ApiFromEmail", "api_from_email", db.Varchar)
	//info.AddField("ApiFromName", "api_from_name", db.Varchar)
	info.AddField("SendingMethod", "sending_method", db.Varchar)
	info.AddField("Основной сервер", "main_server", db.Boolean)
	info.AddField("Сервер активен", "is_active", db.Boolean)
	//info.AddField("PreferredDomains", "preferred_domains", db.Int)

	info.SetTable("email_servers").SetTitle("Почтовые сервера").SetDescription("Почтовые сервера")

	formList := emailServersTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("EmailDefaultFrom", "email_default_from", db.Varchar, form.Text)
	formList.AddField("EmailHost", "email_host", db.Varchar, form.Text)
	formList.AddField("EmailPort", "email_port", db.Int, form.Text)
	formList.AddField("EmailUsername", "email_username", db.Varchar, form.Text)
	formList.AddField("SSL", "email_use_ssl", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	formList.AddField("TLS", "email_use_tls", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	formList.AddField("EmailFailSilently", "email_fail_silently", db.Varchar, form.Text)
	formList.AddField("EmailTimeout", "email_timeout", db.Int, form.Number).FieldDefault("0")
	formList.AddField("EmailSSLCertFile", "email_ssl_cert_file", db.Varchar, form.File)
	formList.AddField("EmailSSLKeyfile", "email_ssl_key_file", db.Varchar, form.File)
	formList.AddField("ApiKey", "api_key", db.Varchar, form.Text)
	formList.AddField("ApiUsername", "api_username", db.Varchar, form.Text)
	formList.AddField("ApiFromEmail", "api_from_email", db.Varchar, form.Text)
	formList.AddField("ApiFromName", "api_from_name", db.Varchar, form.Text)
	formList.AddField("SendingMethod", "sending_method", db.Varchar, form.Text)
	formList.AddField("Основной сервер", "main_server", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	formList.AddField("Сервер активен", "is_active", db.Boolean, form.Switch).
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	formList.AddField("PreferredDomains", "preferred_domains", db.Tinyint, form.SelectSingle)

	formList.SetTable("email_servers").SetTitle("Почтовый сервер").SetDescription("Почтовый сервер")

	return
}
