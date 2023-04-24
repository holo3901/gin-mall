package controllers

import (
	"clms/logic"
	"clms/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func AddCarts(ctx *gin.Context) {
	p := new(models.ParamCarts)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddCarts invalid param error", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	id, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	err = logic.AddCarts(p, id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, "success")
}

func GetCarts(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, parseInt := strconv.ParseInt(idStr, 10, 64)
	if parseInt != nil {
		return
	}
	carts, err := logic.GetCarts(id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, carts)
}

func UpdateCarts(ctx *gin.Context) {
	p := new(models.ParamCarts)

	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateCarts invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	err := logic.UpdateCarts(p, id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, "success")

}

func DeleteCarts(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := logic.DeleteCarts(id)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, "success")
}
