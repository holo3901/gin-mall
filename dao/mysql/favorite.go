package mysql

import "clms/models"

func ShowFavorite(p *models.Page, uId string) (favorites []*models.Favorite, total int64, err error) {
	err = dbs.Model(&models.Favorite{}).Preload("User").
		Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}
	// 分页
	err = dbs.Model(models.Favorite{}).Preload("User").Where("user_id=?", uId).
		Offset(int((p.PageNum - 1) * p.PageSize)).
		Limit(int(p.PageSize)).Find(&favorites).Error
	return
}

func AddFavorite(favorite *models.Favorite) (err error) {
	err = dbs.Model(&models.Favorite{}).Create(favorite).Error
	return
}

func GetFavoriteByIdS(product int64, username int64) (count int64, err error) {
	err = dbs.Model(&models.Favorite{}).Where("id=? AND user_id=?", product, username).Count(&count).Error
	return
}

func GetFavoritesByIds(list []string) (product []*models.Favorite, err error) {
	err = dbs.Model(&models.Favorite{}).Find(&product, list).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetFavoriteById(id int64) (favorite *models.Favorite, err error) {
	err = dbs.Model(&models.Favorite{}).Where("id=?", id).Find(&favorite).Error
	return
}

func GetFavoriteByIdm(id int64) (favorite []*models.Favorite, err error) {
	err = dbs.Model(&models.Favorite{}).Where("user_id=?", id).Find(&favorite).Error
	return
}

func DeleteFavorite(id int64) error {
	err := dbs.Model(&models.Favorite{}).Delete(&models.Favorite{}, id).Error
	return err
}
