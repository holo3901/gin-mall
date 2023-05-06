package mysql

import (
	"clms/models"
	"database/sql"
)

func ProductList() (data []*models.Product, err error) {
	err = dbs.Model(&models.Product{}).Find(&data).Error
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return
}

func GetProductsByIds(list []string) (product []*models.Product, err error) {
	err = dbs.Model(&models.Product{}).Find(&product, list).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetProductById(id int64) (product *models.Product, err error) {
	err = dbs.Model(&models.Product{}).Where("id=?", id).Find(&product).Error
	return
}
func SearchProduct(p *models.ParamProduct) (products []*models.Product, err error) {
	err = dbs.Model(&models.Product{}).Where("name LIKE ? OR info LIKE ?", "%"+p.Info+"%", "%"+p.Info+"%").
		Offset(int((p.PageNum - 1) * p.PageSize)).Limit(int(p.PageSize)).Find(&products).Error
	return
}

func CreateProductImg(p *models.ProductImg) error {
	err := dbs.Model(&models.ProductImg{}).Create(&p).Error
	return err
}
func ListProductImg(p int) (img []*models.ProductImg, err error) {
	err = dbs.Model(&models.ProductImg{}).Where("id=?", p).Find(&img).Error
	return
}

func GetProductsById(id int64) (products []*string, err error) {
	err = dbs.Model(&models.Product{}).Select("id").Where("category_id=?", id).Find(&products).Error
	return
}

func CreateProduct(product *models.Product) error {
	return dbs.Model(&models.Product{}).Create(&product).Error
}

func UpdateProduct(product *models.Product, id string) (err error) {
	err = dbs.Model(&models.Product{}).Where("id =?", id).Updates(&product).Error
	return
}

func DeleteProduct(id int64) error {
	err := dbs.Model(&models.Product{}).Delete(&models.Product{}, id).Error
	return err
}
