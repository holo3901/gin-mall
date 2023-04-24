package logic

import (
	"clms/dao/mysql"
	"clms/models"
	"clms/pkg/JWT"
	"clms/pkg/encryption"
	"clms/settings"
	"database/sql"
	"fmt"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
)

func UserRegister(user *models.ParamRegister) error {
	//1.判断是否被注册
	_, err := mysql.GetUserById(user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	users := &models.User{
		UserName:       user.UserName,
		PasswordDigest: user.Password,
		NickName:       user.NickName,
		Status:         models.Active,
		Money:          encryption.Encrypt.AesEncoding("10000"),
	}
	//2.对用户密码进行加密,并存入数据库
	fmt.Println("user")
	return mysql.Register(users)
}

func UserLogin(user *models.ParamLogin) (users *models.User, token string, err error) {
	fmt.Println(user)
	users = &models.User{
		UserName:       user.UserName,
		PasswordDigest: user.PassWord,
	}
	if users, err = mysql.Login(users); err != nil {
		return nil, "", err
	}
	token, err = JWT.GenToken(int64(users.ID), users.UserName)
	if err != nil {
		return nil, "", err
	}
	return
}

func UserUpdate(user *models.ParamUpdateUser, username int64) (token string, err error) {
	v := settings.EmailConfig{}
	users, err := mysql.GetUserByIds(uint(username))
	if err != nil {
		return "", err
	}
	if user.NickName != "" {
		users.NickName = user.NickName
		err = mysql.UserUpDate(username, users)
		if err != nil {
			return "", err
		}
	}
	if user.OperationType != 0 {
		token, err = JWT.GenerateEmailToken(username, user.OperationType, user.Email, user.PassWord)
		if err != nil {
			return "", err
		}
		a := models.Notice{
			Text: token,
		}
		err = mysql.CreateNotice(&a)
		if err != nil {
			return "", err
		}
		address := v.ValidEmail + token
		mailStr := token
		mailText := strings.Replace(mailStr, "Email", address, -1)
		m := mail.NewMessage()
		m.SetHeader("From", v.SmtpEmail)
		m.SetHeader("To", user.Email)
		m.SetHeader("Subject", "holo")
		m.SetBody("text/html", mailText)
		d := mail.NewDialer(v.SmtpHost, 465, v.SmtpEmail, v.SmtpPass)
		d.StartTLSPolicy = mail.MandatoryStartTLS
		if err = d.DialAndSend(m); err != nil {
			return "", err
		}
	}
	return
}

func UserValid(valid *models.ParamUserValid) error {
	info, err := JWT.ParseEmailToken(valid.Token)
	if err != nil {
		return err
	}
	users, err := mysql.GetUserByIds(uint(info.UserName))

	if info.OperationType == 1 {
		users.Email = info.Email
		err = mysql.UserUpDate(info.UserName, users)
		if err != nil {
			return err
		}
	} else if info.OperationType == 2 {
		users.Email = ""
		err = mysql.UserUpDate(info.UserName, users)
		if err != nil {
			return err
		}
	} else {
		users.PasswordDigest = info.Password
		err = mysql.UserUpDate(info.UserName, users)
		if err != nil {
			return err
		}

	}
	return nil
}

func Post(file multipart.File, fileSize int64, username int64) error {
	path, err := UploadToQiNiu(file, fileSize)
	if err != nil {
		return err
	}
	user, err := mysql.GetUserByIds(uint(username))
	if err != nil {
		return err
	}
	user.Avatar = path
	err = mysql.UserUpDate(username, user)
	if err != nil {
		return err
	}
	return nil
}
