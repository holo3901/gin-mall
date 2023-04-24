package controllers

import (
	"clms/logic"
	"clms/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func ShowFavorite(ctx *gin.Context) {
	pageSize := ctx.Param("pagesize")
	size, err := strconv.ParseInt(pageSize, 10, 64)
	pagenum := ctx.Param("pagenum")
	page, err := strconv.ParseInt(pagenum, 10, 64)
	a := &models.Page{
		PageSize: size,
		PageNum:  page,
	}
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	username, err := GetCurrentUserID(ctx)

	favorite, i, err := logic.ShowFavorite(a, username)
	if err != nil {
		return
	}
	ResponseSuccess(ctx, gin.H{
		"data":  favorite,
		"total": i,
	})
}

func AddFavorite(ctx *gin.Context) {
	p := new(models.ParamFavorite)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("add favorite valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	username, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	err = logic.AddFavorite(p, username)
	if err != nil {
		zap.L().Error("添加喜欢物品失败,", zap.Error(err))
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	ResponseSuccess(ctx, CodeSuccess)
}

func DeleteFavorite(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //"10"表示是十进制,bitSize表示是Int64
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	err = logic.DeleteFavorite(id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, CodeSuccess)
}
