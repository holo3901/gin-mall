package mysql

import "clms/models"

func ListCarousels() (data []*models.Carousel, err error) {
	err = dbs.Model(&models.Carousel{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return
}
