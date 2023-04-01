package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Int))

	info := users.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("名称", "name", db.Varchar)
	info.AddField("性别", "gender", db.Tinyint)
	info.AddField("City", "city", db.Varchar)
	info.AddField("Ip", "ip", db.Varchar)
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("users").SetTitle("用户").SetDescription("用户")

	formList := users.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("名称", "name", db.Varchar, form.Text)
	formList.AddField("性别", "gender", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "未知", Value: "0"},
			{Text: "男", Value: "1"},
			{Text: "女", Value: "2"},
		}).FieldDefault("0")
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Ip", "ip", db.Varchar, form.Ip)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("users").SetTitle("用户").SetDescription("用户")

	return users
}
