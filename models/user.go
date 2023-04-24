package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}

const (
	Active string = "active" //激活用户
)
