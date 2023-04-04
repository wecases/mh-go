package models

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (User) TableName() string {
	return "users"
}

type User struct {
	Model
	ParentID   int64         `gorm:"comment:'父节点id'" json:"parent_id"`
	InviteCode string        `gorm:"comment:'邀请码'" json:"invite_code"`
	Path       pq.Int64Array `gorm:"type:text;null;comment:'路径'" json:"-"`

	Name     string `gorm:"not null;comment:'用户名'" json:"name"`
	Phone    string `gorm:"unique;uniqueIndex;not null;comment:'手机号'" json:"phone"`
	Password string `gorm:"not null;comment:'密 码'" json:"-"`
	Avatar   string `gorm:"comment:'头像'" json:"avatar"`
}

// 脱敏处理
// func (u *User) Sanitize() map[string]interface{} {
// 	return map[string]interface{}{
// 		"name":  u.Name,
// 		"phone": u.Phone[:3] + "****" + u.Phone[7:],
// 	}
// }

// 加密密码
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// 校验密码
func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
