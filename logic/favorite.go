package logic

import (
	"clms/dao/mysql"
	"clms/dao/redis"
	"clms/models"
	"database/sql"
	"strconv"
	"time"
)

func ShowFavorite(a *models.Page, username int64) (favorites []*models.Favorite, total int64, err error) {
	order, err := redis.GetFavoriteOrder(a, username)
	if err != nil {
		return nil, 0, err
	}
	ids, err := mysql.GetFavoritesByIds(order)
	if err != nil {
		return nil, 0, err
	}
	return ids, int64(len(ids)), nil

}

func AddFavorite(a *models.ParamFavorite, username int64) error {

	//1,查询是否存在
	users, err := mysql.GetUserByIds(uint(username))
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	s, err := mysql.GetFavoriteByIdS(a.ProductId, int64(users.ID))
	if err != nil {
		return err
	}
	if s != 0 {
		return err
	}

	products, err := mysql.GetProductById(a.ProductId)
	if err != nil {
		return err
	}
	boss, err := mysql.GetUserById(strconv.FormatInt(a.BossId, 10))
	if err != nil {
		return err
	}
	favorite := models.Favorite{
		User:      *users,
		UserID:    users.ID,
		Product:   *products,
		ProductID: uint(a.ProductId),
		Boss:      *boss,
		BossID:    boss.ID,
	}
	favorite.ID = uint(time.Now().Unix())
	err = mysql.AddFavorite(&favorite)
	if err != nil {
		return err
	}
	err = redis.CreateFavorite(username, int(favorite.ID))
	if err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(id int64) error {
	byId, err := mysql.GetFavoriteById(id)
	if err != nil {
		return err
	}
	err = mysql.DeleteFavorite(id)
	if err != nil {
		return err
	}
	err = redis.RefreshFavoriteOrder(byId.User.UserName, int64(byId.UserID))
	if err != nil {
		return err
	}
	return nil
}
