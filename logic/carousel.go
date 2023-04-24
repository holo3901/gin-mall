package logic

import (
	"clms/dao/mysql"
	"clms/models"
	"go.uber.org/zap"
)

func ListCarousels() ([]*models.Carousel, error) {
	carousels, err := mysql.ListCarousels()
	if err != nil {
		zap.L().Error("mysql.ListCarousels default", zap.Error(err))
		return nil, err
	}
	return carousels, nil
}
