package logic

import (
	"clms/dao/mysql"
	"clms/dao/redis"
	"clms/models"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

func ProductList() ([]*models.Product, error) {
	return mysql.ProductList()
}

func ProductListByID(id int64) ([]*models.Product, error) {
	order, err := redis.GetProductsOrder(id)
	if err != nil {
		return nil, err
	}
	if len(order) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return nil, err
	}
	zap.L().Debug("GetProductsOrder", zap.Any("ids", order))
	return mysql.GetProductsByIds(order)
}

func SearchProduct(p *models.ParamProduct) ([]*models.Product, error) {
	return mysql.SearchProduct(p)
}

func ListProductImg(id int64) ([]*models.ProductImg, error) {
	img, err := mysql.ListProductImg(int(id))
	if err != nil {
		return nil, err
	}
	var a []*models.ProductImg
	for _, v := range img {
		a = append(a, &models.ProductImg{
			ProductID: v.ProductID,
			ImgPath:   v.ImgPath,
		})
	}
	return a, nil
}

func CreateProduct(p *models.ParamProductService, file []*multipart.FileHeader, username int64) error {

	boss, _ := mysql.GetUserByIds(uint(username))
	tmp, _ := file[0].Open()
	path, err := UploadToQiNiu(tmp, file[0].Size)
	if err != nil {
		return err
	}
	product := &models.Product{
		Name:          p.Name,
		CategoryID:    uint(p.CategoryID),
		Title:         p.Title,
		Info:          p.Info,
		ImgPath:       path,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		OnSale:        true,
		BossID:        int(boss.ID),
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	product.ID = uint(time.Now().Unix())
	err = mysql.CreateProduct(product)
	if err != nil {
		return err
	}
	err = redis.CreateProductOrder(int64(product.CategoryID), int64(product.ID))
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(p *models.ParamProductService, file []*multipart.FileHeader, id string) error {
	tmp, _ := file[0].Open()
	path, err := UploadToQiNiu(tmp, file[0].Size)
	product := &models.Product{
		Name:          p.Name,
		CategoryID:    uint(p.CategoryID),
		Title:         p.Title,
		Info:          p.Info,
		ImgPath:       path,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		OnSale:        p.OnSale,
	}

	err = mysql.UpdateProduct(product, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int64) error {
	byId, err := mysql.GetProductById(id)
	if err != nil {
		return err
	}
	err = mysql.DeleteProduct(id)
	if err != nil {
		return err
	}
	err = redis.RefreshProductOrder(int64(byId.CategoryID))
	if err != nil {
		return err
	}
	return nil
}
