package models

import (
	"fmt"
	"reflect"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
	DB  = orm
)

func Init(c db.Connection) {
	orm, err = gorm.Open("mysql", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}

	// 执行迁移
	migrator()
}

// 迁移数据表
func migrator() {
	fmt.Println("开始迁移表")

	for _, table := range Tables {
		// orm.DropTable(table) // 删除表

		result := orm.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(table)

		if result.Error != nil {
			panic(result.Error)
		}

		fmt.Println("迁移了", reflect.TypeOf(table).Elem().Name())
	}

	fmt.Println("迁移结束")
}
