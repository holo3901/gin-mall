package logic

import (
	"clms/dao/mysql"
	"clms/pkg/encryption"
	"fmt"
	"strconv"
)

func OrderPay(orderId int64, username int64) (err error) {

	user, err := mysql.GetUserByIds(uint(username))
	if err != nil {
		return
	}
	id, err := mysql.GetOrderById(orderId)
	if err != nil {
		return
	}
	money := id.Money
	num := id.Num
	money = money * float64(num)
	moneyStr := encryption.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	if moneyFloat-money < 0.0 { // 金额不足进行回滚
		return
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = encryption.Encrypt.AesEncoding(finMoney)
	err = mysql.UserUpDate(username, user)

	boss, err := mysql.GetUserByIds(id.BossID)
	if err != nil {
		return
	}
	m := encryption.Encrypt.AesDecoding(boss.Money)
	bossmoneyFloat, _ := strconv.ParseFloat(m, 64)
	find := fmt.Sprintf("%f", bossmoneyFloat+money)
	boss.Money = encryption.Encrypt.AesEncoding(find)

	err = mysql.UserUpDate(username, boss)

	byId, err := mysql.GetProductById(int64(id.ProductID))
	if err != nil {
		return
	}
	byId.Num = byId.Num - num
	err = mysql.UpdateProduct(byId, strconv.Itoa(int(byId.ID)))
	if err != nil {
		return
	}
	return
}
