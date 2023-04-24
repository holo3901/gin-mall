package mysql

import "clms/models"

func ListCategory() (data []*models.Category, err error) {
	err = dbs.Model(&models.Category{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return
}
