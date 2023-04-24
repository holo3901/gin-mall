package controllers

import (
	"clms/logic"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func ListCarousels(ctx *gin.Context) {
	carousels, err := logic.ListCarousels()
	if err != nil {
		if err == sql.ErrNoRows {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(ctx, carousels)
}
