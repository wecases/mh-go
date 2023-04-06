package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFilemanagerSettingTable(ctx *context.Context) table.Table {

	filemanagerSetting := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := filemanagerSetting.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Key", "key", db.Varchar)
	info.AddField("Value", "value", db.Text)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("filemanager_setting").SetTitle("FilemanagerSetting").SetDescription("FilemanagerSetting")

	formList := filemanagerSetting.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Key", "key", db.Varchar, form.Text)
	formList.AddField("Value", "value", db.Text, form.RichText)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("filemanager_setting").SetTitle("FilemanagerSetting").SetDescription("FilemanagerSetting")

	return filemanagerSetting
}
