package controllers

import (
	"clms/logic"
	"github.com/gin-gonic/gin"
	"strconv"
)

func OrderPay(ctx *gin.Context) {
	idStr := ctx.Param("id")
	OrderId, _ := strconv.ParseInt(idStr, 10, 64)
	userID, err := GetCurrentUserID(ctx)
	if err != nil {
		return
	}
	err = logic.OrderPay(OrderId, userID)
	if err != nil {
		ResponseError(ctx,CodeServerBusy)
		return
	}
	ResponseSuccess(ctx,"success")
}
