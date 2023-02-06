package responsex

import (
	"context"
	"fmt"
	"gin-skeleton/common/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceid"`
}

func NoMethod(ctx *gin.Context) {
	Result(ctx, errorx.NewCodeError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)))
}

func NoRoute(ctx *gin.Context) {
	Result(ctx, errorx.NewCodeError(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
}

func ServiceUnavailable(ctx *gin.Context) {
	Result(ctx, errorx.NewCodeError(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable)))
}

func Result(ctx *gin.Context, data interface{}) {
	result := newResponse(ctx, data)
	result.TraceId = fmt.Sprintf("%s", ctx.Value("traceId"))
	httpStatus := http.StatusOK
	if result.Code >= 100 && result.Code <= 511 {
		httpStatus = result.Code
	}
	ctx.JSON(httpStatus, result)
}

func newResponse(ctx context.Context, data interface{}) *response {
	switch value := data.(type) {

	case *errorx.CodeError:
		return &response{
			Code:    value.Code,
			Message: value.Msg,
		}
	case error:
		return &response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return &response{
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}
