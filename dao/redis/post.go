package redis

import (
	"clms/dao/mysql"
	"clms/models"
	"go.uber.org/zap"
	"strconv"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	//3.ZREVRANGE按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
func CreatePost(postID int64, communityID int64) error {

	//把帖子id加到社区的set
	cKey := getRedisKey(keyCommunitySetPF + strconv.Itoa(int(communityID)))
	if _, err := client.SAdd(cKey, postID).Result(); err != nil {
		zap.L().Error("getRedisKey default", zap.Error(err))
		return err
	}
	return nil
}
func GetProductsOrder(category int64) ([]string, error) {
	//从redis获取id
	//1.根据用户请求中携带的order参数确定要查询的redis key
	orderKey := getRedisKey(keyCategorySetPF + strconv.Itoa(int(category)))
	//2.确定查询的索引起始点
	return getIDsFormKey(orderKey, 0, 0)
}

func CreateProductOrder(category int64, product int64) error {
	ckey := getRedisKey(keyCategorySetPF + strconv.Itoa(int(category)))
	if _, err := client.SAdd(ckey, product).Result(); err != nil {
		zap.L().Error("getRedisKey default", zap.Error(err))
		return err
	}
	return nil
}

func RefreshProductOrder(category int64) error {
	skey := getRedisKey(keyCategorySetPF + strconv.Itoa(int(category)))
	_, err := client.Del(skey).Result()
	if err != nil {
		zap.L().Error("delRedisKey default", zap.Error(err))
		return err
	}
	ids, err2 := mysql.GetProductsById(category)
	if err2 != nil {
		return err2
	}
	for _, v := range ids {
		if _, err = client.SAdd(skey, v).Result(); err != nil {
			zap.L().Error("getRedisKey default", zap.Error(err))
			return err
		}
	}
	return nil
}

func CreateFavorite(user int64, favorite int) error {
	ckey := getRedisKey(KeyFavoriteSetPF + strconv.Itoa(int(user)))
	if _, err := client.SAdd(ckey, favorite).Result(); err != nil {
		zap.L().Error("getrediskey default", zap.Error(err))
		return err
	}
	return nil
}

func GetFavoriteOrder(a *models.Page, user int64) ([]string, error) {
	orderKey := getRedisKey(KeyFavoriteSetPF + strconv.Itoa(int(user)))
	//2.确定查询的索引起始点
	return getIDsFormKey(orderKey, a.PageNum, a.PageSize)
}
func RefreshFavoriteOrder(username string, id int64) error {
	skey := getRedisKey(keyCategorySetPF + username)
	_, err := client.Del(skey).Result()
	if err != nil {
		zap.L().Error("delRedisKey default", zap.Error(err))
		return err
	}
	ids, err2 := mysql.GetFavoriteByIdm(id)
	if err2 != nil {
		return err2
	}
	for _, v := range ids {
		if _, err = client.SAdd(skey, v).Result(); err != nil {
			zap.L().Error("getRedisKey default", zap.Error(err))
			return err
		}
	}
	return nil
}
