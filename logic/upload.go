package logic

import (
	"clms/settings"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadToQiNiu(file multipart.File, fileSize int64) (path string, err error) {
	a := settings.Conf.QiMiuConfig
	var AccessKey = a.AccessKey
	var SerectKey = a.SerectKey
	var Bucket = a.Bucket
	var ImgUrl = a.QiniuServe
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	fmt.Println(a)
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
        fmt.Println(err)
		return "", err
	}
	url := ImgUrl + ret.Key
	return url, nil
}

