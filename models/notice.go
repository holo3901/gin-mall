package models

import "gorm.io/gorm"

// Notice 公告模型 存放公告和邮件模板
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
