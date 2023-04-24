package mysql

import "clms/models"

// CreateNotice 创建notice
func CreateNotice(notice *models.Notice) error {
	return dbs.Model(&models.Notice{}).Create(&notice).Error
}
