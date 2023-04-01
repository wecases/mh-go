package models

// type User struct {
// 	gorm.Model
// 	ID         uint `gorm:"primary_key,column:cpu"`
// 	CPU        uint `gorm:"column:cpu"`
// 	Likes      uint `gorm:"column:likes"`
// 	Sales      uint `gorm:"column:sales"`
// 	NewMembers uint `gorm:"column:new_members"`

// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

//	func GetUser(where interface{}) *User {
//		user := new(User)
//		if err := orm.Model(&user).Where(where).First(user).Error; err != nil {
//			panic(err)
//		}
//		return user
//	}

func GetUser(where ...interface{}) (user *User, err error) {
	user = &User{}

	if err := orm.Model(&user).First(&user, where...).Error; err != nil {
		panic(err)
	}

	return user, err
}

func GetUserList() *User {
	user := new(User)
	if err := orm.Model(&user).Select("*").Error; err != nil {
		panic(err)
	}
	return user
}
