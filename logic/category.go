package logic

import (
	"clms/dao/mysql"
	"clms/models"
	"go.uber.org/zap"
)

func ListCategory() ([]*models.Category, error) {
	category, err := mysql.ListCategory()
	if err != nil {
		zap.L().Error("mysql.ListCategory default", zap.Error(err))
		return nil, err
	}
	return category, nil
}
