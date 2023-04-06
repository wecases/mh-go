package tables

import (
	"mh-go/models"
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql").SetPrimaryKey("id", db.Bigint))

	info := users.GetInfo().HideFilterArea().HideDeleteButton()

	info.AddField("序号", "id", db.Bigint).FieldFilterable().FieldSortable()
	info.AddField("名字", "name", db.Varchar).FieldFilterable()
	info.AddField("手机号", "phone", db.Varchar).FieldFilterable()
	info.AddField("头像", "avatar", db.Varchar).FieldImage("50", "50")
	info.AddField("上级", "parent_id", db.Bigint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" || model.Value == "" || model.Value == "null" || model.Value == "NULL" {
			return ""
		}
		var user models.User
		models.DB.Model(user).Where("id = ?", model.Value).First(&user)
		return user.Phone
	})
	info.AddField("邀请码", "invite_code", db.Varchar).FieldFilterable()
	info.AddField("团队人数", "path", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {

		// fmt.Println("path", model.Value)
		// if model.Row["path"] != nil {
		// 	slice := make([]string, 0)
		// 	slice = append(slice, strings.Split(model.Row["path"].(string), ",")...)
		// 	fmt.Println("slice", slice)
		// 	// fmt.Printf("slice type %T", slice)
		// 	var user1 []models.User
		// 	models.DB.Where("invite_code IN (?)", slice).Find(&user1)
		// 	fmt.Println("user1", user1)
		// }

		invite_code := model.Row["invite_code"].(string)

		var user []models.User
		models.DB.Where("path LIKE ?", "%"+invite_code+"%").Find(&user)

		return strconv.Itoa(len(user))
	})
	info.AddField("创建时间", "created_at", db.Datetime).FieldSortable().FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("更新时间", "updated_at", db.Datetime).FieldSortable().FieldFilterable(types.FilterType{FormType: form.DatetimeRange})

	info.SetTable("users").SetTitle("用户").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("序号", "id", db.Bigint, form.Default).FieldHide()
	formList.AddField("名字", "name", db.Varchar, form.Text)
	formList.AddField("手机号", "phone", db.Varchar, form.Text)
	formList.AddField("头像", "avatar", db.Varchar, form.File)
	formList.AddField("上级", "parent_id", db.Bigint, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "man", Value: "0"},
		{Text: "women", Value: "1"},
	})
	formList.AddField("邀请码", "invite_code", db.Varchar, form.Text)

	formList.SetTable("users").SetTitle("用户").SetDescription("Users")

	return users
}
