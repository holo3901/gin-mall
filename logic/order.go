package logic

import (
	"clms/dao/mysql"
	"clms/models"
	"time"
)

func AddOrder(server *models.ParamOrder, username int64) error {
	id, err := mysql.GetUserByIds(uint(username))
	if err != nil {
		return err
	}
	orders, err := mysql.GetAddressById(int64(server.AddressID))
	if err != nil {
		return err
	}
	order := models.Order{
		UserID:    id.ID,
		ProductID: server.ProductID,
		BossID:    server.BossID,
		AddressID: orders.ID,
		Num:       int(server.Num),
		OrderNum:  uint64(server.Num),
		Type:      1,
		Money:     float64(server.Money),
	}
	order.OrderNum = uint64(time.Now().Unix())
	order.ID = uint(time.Now().Unix())
	err = mysql.CreateOrder(&order)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(username int64) (orders []*models.Order, num int64, err error) {
	id, err := mysql.GetUserByIds(uint(username))
	if err != nil {
		return nil, 0, err
	}

	ids, err := mysql.GetOrdersByIds(int64(id.ID))
	if err != nil {
		return nil, 0, err
	}

	return ids, int64(len(ids)), nil
}

func GetOrderById(id int64) (order *models.Order, err error) {
	return mysql.GetOrderById(id)
}

func DeleteOrder(id int64) error {

	err := mysql.DeleteOrder(id)
	return err
}
