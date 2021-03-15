package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetEmailsTable return the model of table email server.
func GetEmailsTable(ctx *context.Context) (emailsTable table.Table) {

	emailsTable = table.NewDefaultTable(table.Config{
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

	info := emailsTable.GetInfo()
	info.AddField("SUBJECT", "subject", db.Varchar)
	info.AddField("EMAIL BODY", "body", db.Varchar)
	info.AddField("SENDER", "sender", db.Varchar)
	info.AddField("RECIPIENT", "recipient", db.Varchar)
	info.AddField("NEWSLETTER", "newsletter", db.Varchar)
	info.AddField("STATUS", "status", db.Varchar)
	info.AddField("TYPE", "type", db.Varchar)
	//info.AddField("CREATEDATATIME", "body", db.Varchar)
	//info.AddField("CHANGEDATATIME", "body", db.Varchar)

	info.SetTable("emails").SetTitle("Письма").SetDescription("Письма")

	formList := emailsTable.GetForm()
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Sender", "sender", db.Varchar, form.Text)
	formList.AddField("Recipient", "recipient", db.Varchar, form.Text)
	formList.AddField("Body", "body", db.Varchar, form.Text)
	formList.AddField("Subject", "subject", db.Varchar, form.Text)
	formList.AddField("Newsletter", "newsletter", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Text)
	formList.AddField("Хеш статуса для безгеморройного индексирования", "StatusHash", db.Varchar, form.Text)
	formList.AddField("Email remote id", "email_remote_id", db.Int, form.Text)
	formList.AddField("Used server", "used_server", db.Int, form.Text)
	formList.AddField("Type", "email_default_from", db.Varchar, form.Text)

	formList.SetTable("emails").SetTitle("Письмо").SetDescription("Письмо")

	return
}
