package mysql

import "clms/models"

func CreateCarts(p *models.Cart) (err error) {
	err = dbs.Model(&models.Cart{}).Create(&p).Error
	return
}

func GetCartsById(pro uint, boss uint, name uint) (cart *models.Cart, err error) {
	err = dbs.Model(&models.Cart{}).Where("product_id=? AND boss_id AND user_id =?", pro, boss, name).Find(&cart).Error
	return
}

func SaveCarts(p *models.Cart) (err error) {
	err = dbs.Model(&models.Cart{}).Save(&p).Error
	return
}
func GetCarts(p int64) (x *models.Cart, err error) {
	err = dbs.Model(&models.Cart{}).Where("id=?", p).Find(&x).Error
	return
}

func DeleteCarts(p int64) error {
	err := dbs.Model(&models.Cart{}).Delete(&models.Cart{}, p).Error
	return err
}
