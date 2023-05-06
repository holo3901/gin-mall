package mysql

import "clms/models"

// CreateNotice 创建notice
func CreateNotice(notice *models.Notice) error {
	return dbs.Model(&models.Notice{}).Create(&notice).Error
}
func GetNoticeByIds(id uint) (notice *models.Notice, err error) {
	err = dbs.Model(&models.Notice{}).Where("id =?", id).Find(&notice).Error
	return
}
