package mysql

import "clms/models"

func GetAddressById(id int64) (order *models.Address, err error) {
	err = dbs.Model(&models.Address{}).Where("id =?", id).Find(order).Error
	return
}
