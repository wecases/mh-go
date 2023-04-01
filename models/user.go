package models

func (User) TableName() string {
	return "users"
}

type User struct {
	Model
	Name     string `gorm:"not null;comment:'用户名'" json:"name"`
	Phone    string `gorm:"not null;comment:'手机号'" json:"phone"`
	Password string `gorm:"not null;comment:'密 码'" json:"-"`
}

// 根据条件查询单条数据
func GetUser(where interface{}) (*User, error) {
	user := new(User)
	if err := orm.Model(&user).Where(where).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserList() []*User {
	var users []*User
	if err := orm.Find(&users).Error; err != nil {
		panic(err)
	}
	return users
}

// 脱敏处理
func (u *User) Sanitize() map[string]interface{} {
	return map[string]interface{}{
		"name":  u.Name,
		"phone": u.Phone[:3] + "****" + u.Phone[7:],
	}
}
