package mysql

import (
	"clms/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const secret = "miku"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword))) //EncodeToString把字节转化为16进制的字符串
}

func GetUserById(name string) (user *models.User, err error) {
	err = dbs.Model(&models.User{}).Where("user_name=?", name).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByIds(name uint) (user *models.User, err error) {
	err = dbs.Model(&models.User{}).Where("id=?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Register(user *models.User) error {
	user.PasswordDigest = encryptPassword(user.PasswordDigest)
	err := dbs.Model(&models.User{}).Create(&user).Error
	fmt.Println(err)
	return err
}

func Login(user *models.User) (a *models.User, err error) {
	err = dbs.Model(&models.User{}).Where("user_name=?", user.UserName).First(&a).Error
	if err != nil {
		return nil, err
	}
	user.PasswordDigest = encryptPassword(user.PasswordDigest)
	if a.PasswordDigest != user.PasswordDigest {
		return nil, err
	}
	return
}

func UserUpDate(username int64, user *models.User) error {
	err := dbs.Model(&models.User{}).Where("id=?", username).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}
