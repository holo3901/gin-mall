package controllers

import (
	"clms/logic"
	"clms/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func ListProduct(ctx *gin.Context) {
	data, err := logic.ProductList()
	if err != nil {
		zap.L().Error("logic.ProductList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}

func ListProductById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	byID, err := logic.ProductListByID(id)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, byID)
}

func SearchProduct(ctx *gin.Context) {
	p := new(models.ParamProduct)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("searchProduct with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	product, err := logic.SearchProduct(p)
	if err != nil {
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, product)
}

func ListProductImg(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	img, err := logic.ListProductImg(id)
	if err != nil {
		zap.L().Error("ListProductImg failed", zap.Error(err))
		if err == sql.ErrNoRows {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, img)
}

func AddProduct(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	files := form.File["file"]
	if err != nil {
		zap.L().Error("AddProduct invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	p := new(models.ParamProductService)
	if err = ctx.ShouldBind(p); err != nil {
		zap.L().Error("AddProduct invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)

		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	username, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	err = logic.CreateProduct(p, files, username)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	ResponseSuccess(ctx, gin.H{
		"info": "创建成功",
	})
}

func UpdateProduct(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	files := form.File["file"]
	p := new(models.ParamProductService)
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("updateProduct invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	err = logic.UpdateProduct(p, files, id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, CodeSuccess)
}

func DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //"10"表示是十进制,bitSize表示是Int64
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	err = logic.DeleteProduct(id)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	ResponseSuccess(ctx, CodeSuccess)
}
