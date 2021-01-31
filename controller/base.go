package controller

import (
	"gin-demo/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    constant.ResponseCode `json:"code"`
	Message string                `json:"message"`
	Data    interface{}           `json:"data"`
}

func setContent(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
}
func Success(ctx *gin.Context, data interface{}) {
	resp := response{
		Code:    constant.SUCCESS,
		Message: constant.GetCodeText(constant.SUCCESS),
		Data:    data,
	}
	setContent(ctx, http.StatusOK, resp)
}

func SuccessWithMessage(ctx *gin.Context, message string, data interface{}) {
	resp := response{
		Code:    constant.SUCCESS,
		Message: message,
		Data:    data,
	}
	setContent(ctx, http.StatusOK, resp)
}

func Error(ctx *gin.Context, responseCode constant.ResponseCode) {
	resp := response{
		Code:    responseCode,
		Message: constant.GetCodeText(responseCode),
		Data:    [0]int{},
	}
	setContent(ctx, http.StatusOK, resp)
}

func ErrorWithMessage(ctx *gin.Context, responseCode constant.ResponseCode, message string) {
	resp := response{
		Code:    responseCode,
		Message: message,
		Data:    [0]int{},
	}
	setContent(ctx, http.StatusOK, resp)
	ctx.Abort()
}

//NOTFOUND 页面不存在
func NOTFOUND(ctx *gin.Context) {
	resp := response{
		Code:    constant.CODE_404,
		Message: constant.GetCodeText(constant.CODE_404),
		Data:    [0]int{},
	}
	setContent(ctx, http.StatusNotFound, resp)
}

//StatuscoreServerError 服务器内部错误
func StatuscoreServerError(ctx *gin.Context) {
	resp := response{
		Code:    constant.CODE_500,
		Message: constant.GetCodeText(constant.CODE_500),
		Data:    [0]int{},
	}
	setContent(ctx, http.StatusInternalServerError, resp)
}
