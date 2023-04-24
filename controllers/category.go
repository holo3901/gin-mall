package controllers

import (
	"clms/logic"
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListCategory(ctx *gin.Context) {
	category, err := logic.ListCategory()
	if err != nil {
		if err == sql.ErrNoRows {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeServerBusy, zap.Error(err))
		return
	}
	ResponseSuccess(ctx, category)
}
