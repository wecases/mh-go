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
	ParentID   int64         `json:"parent_id" gorm:"comment:'父节点id'"`
	InviteCode string        `json:"invite_code" gorm:"comment:'邀请码'"`
	Path       pq.Int64Array `json:"-" gorm:"type:text;null;comment:'路径'"`

	Name     string `json:"name" gorm:"not null;comment:'用户名'"`
	Phone    string `json:"phone" gorm:"unique;uniqueIndex;not null;comment:'手机号'"`
	Password string `json:"-" gorm:"not null;comment:'密 码'"`
	Avatar   string `json:"avatar" gorm:"comment:'头像'"`
}

// 脱敏处理
func (u *User) Sanitize() {
	// 对手机号进行脱敏处理
	if len(u.Phone) > 10 {
		u.Phone = u.Phone[:3] + "****" + u.Phone[7:]
	}
}

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
