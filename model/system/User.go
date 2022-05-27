package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"userName" gorm:"comment:用户登录名"` // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`
	Phone    string `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email    string `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}

func (User) TableName() string {
	return "user"
}
