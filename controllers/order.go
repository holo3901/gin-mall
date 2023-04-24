package controllers

import (
	"clms/logic"
	"clms/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func AddOrders(ctx *gin.Context) {
	c := new(models.ParamOrder)
	if err := ctx.ShouldBindJSON(c); err != nil {
		zap.L().Error("shouldbindjson invalid param", zap.Error(err))
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
	err = logic.AddOrder(c, id)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, CodeSuccess)
}

func GetOrders(ctx *gin.Context) {
	username, _ := GetCurrentUserID(ctx)
	orders, i, err := logic.GetOrders(username)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, gin.H{
		"data": orders,
		"num":  i,
	})
}
func GetOrderById(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		zap.L().Error("parseint invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	byId, err := logic.GetOrderById(id)
	if err != nil {
		return
	}
	ResponseSuccess(ctx, byId)

}

func DeleteOrder(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		zap.L().Error("parseint invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	err = logic.DeleteOrder(id)
    if err!=nil{
		ResponseError(ctx,CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx,CodeSuccess)
}
