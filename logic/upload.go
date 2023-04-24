package logic

import (
	"clms/settings"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UploadToQiNiu(file multipart.File, fileSize int64) (path string, err error) {
	a := settings.QiMiuConfig{}
	var AccessKey = a.AccessKey
	var SerectKey = a.SerectKey
	var Bucket = a.Bucket
	var ImgUrl = a.QiniuServe
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
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
		return "", err
	}
	url := ImgUrl + ret.Key
	return url, nil
}
