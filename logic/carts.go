package logic

import (
	"clms/dao/mysql"
	"clms/models"
	"gorm.io/gorm"
)

func AddCarts(p *models.ParamCarts, name int64) error {
	_, err := mysql.GetProductById(int64(p.ProductID))
	if err != nil {
		return err
	}
	byId, err := mysql.GetUserByIds(uint(name))
	if err != nil {
		return err
	}

	cart, err := mysql.GetCartsById(p.ProductID, p.BossID, byId.ID)
	if err == gorm.ErrRecordNotFound {
		cart = &models.Cart{
			UserID:    byId.ID,
			ProductID: p.ProductID,
			BossID:    p.BossID,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}
		err = mysql.CreateCarts(cart)
		if err != nil {
			return err
		}
	} else if cart.Num < cart.MaxNum {
		cart.Num++
		err = mysql.SaveCarts(cart)
		if err != nil {
			return err
		}

	} else {
		return err
	}
	return nil
}

func GetCarts(id int64) (p *models.Cart, err error) {
	p, err = mysql.GetCarts(id)
	if err != nil {
		return nil, err
	}
	return
}

func UpdateCarts(p *models.ParamCarts, id int64) (err error) {
	carts, err := mysql.GetCarts(id)
	if err != nil {
		return err
	}
	cart := models.Cart{
		UserID:    carts.UserID,
		ProductID: carts.ProductID,
		BossID:    carts.BossID,
		Num:       p.Num,
		MaxNum:    carts.MaxNum,
		Check:     false,
	}
	err = mysql.SaveCarts(&cart)
	if err != nil {
		return err
	}
	return
}

func DeleteCarts(id int64) (err error) {
	err = mysql.DeleteCarts(id)
	if err != nil {
		return err
	}
	return
}
