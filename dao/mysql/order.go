package mysql

import "clms/models"

func CreateOrder(m *models.Order) (err error) {
	err = dbs.Model(&models.Order{}).Create(&m).Error
	return
}

func GetOrdersByIds(username int64) (orders []*models.Order, err error) {
	err = dbs.Model(&models.Order{}).Where("id=?", username).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOrderById(id int64) (order *models.Order, err error) {
	err = dbs.Model(&models.Order{}).Where("id=?", id).Find(&order).Error
	return
}

func DeleteOrder(id int64) error {
	err := dbs.Model(&models.Order{}).Delete(&models.Order{}, id).Error
	return err
}
