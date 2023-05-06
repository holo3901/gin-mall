package controllers

import (
	"clms/dao/mysql"
	"clms/logic"
	"clms/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func UserRegister(ctx *gin.Context) {
	p := new(models.ParamRegister)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	if err := logic.UserRegister(p); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		if errors.Is(err, mysql.ErrorUserExist) { //errors.Is判断两个err是否相同
			ResponseError(ctx, CodeUserExist)
			return
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, "success")
}

func UserLogin(ctx *gin.Context) {
	//1. 校验参数
	p := new(models.ParamLogin)
	if err := ctx.ShouldBind(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.登录
	user, token, err := logic.UserLogin(p)
	if err != nil {
		zap.L().Error("userlogin with invalid param", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(ctx, CodeUserNotExist)
			return
		}
		ResponseError(ctx, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, gin.H{
		"data":  user,
		"token": token,
	})
}

func UserUpdate(ctx *gin.Context) {
	p := new(models.ParamUpdateUser)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("update with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	if p.OperationType > 3 || p.OperationType < 0 {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	userid, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	token, err := logic.UserUpdate(p, userid)
	if err != nil {
		zap.L().Error("logic update default", zap.Error(err))
		if err == sql.ErrNoRows {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeServerBusy, err.Error())

		return
	}
	ResponseSuccess(ctx, token)
}

func SendEmail(ctx *gin.Context) {
	l := new(models.ParamSend)
	if err := ctx.ShouldBind(l); err != nil {
		zap.L().Error("sendEmail with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, errs.Translate(trans))
		return
	}
	id, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	err = logic.SendEmail(l, id)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, "success")
}

func ValidEmail(ctx *gin.Context) {
	l := new(models.ParamValid)
	if err := ctx.ShouldBind(l); err != nil {
		zap.L().Error("validEmail with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, errs.Translate(trans))
		return
	}
	err := logic.ValidEmail(l)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, "success")

}
func UpLoadAvatar(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	fileSize := fileHeader.Size
	username, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}

	err = logic.Post(file, fileSize, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, gin.H{
		"info": "头像更新成功",
	})
}
