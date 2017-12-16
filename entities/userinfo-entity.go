package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UserName string `xorm:"notnull pk"` //语义标签
	PassWord string
	Email    string
	CreateAt time.Time `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
