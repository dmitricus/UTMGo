package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetDomainsTable return the model of table email server.
func GetDomainsTable(ctx *context.Context) (domainsTable table.Table) {

	domainsTable = table.NewDefaultTable(table.Config{
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

	info := domainsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Доменное имя", "domain", db.Varchar)

	info.SetTable("domains").SetTitle("Домены").SetDescription("Домены")

	formList := domainsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Доменное имя", "domain", db.Varchar, form.Text)

	formList.SetTable("domains").SetTitle("Домен").SetDescription("Домен")

	return
}
