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
	New(ctx, errorx.NewCodeError(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)))
}

func NoRoute(ctx *gin.Context) {
	New(ctx, errorx.NewCodeError(http.StatusNotFound, http.StatusText(http.StatusNotFound)))
}

func ServiceUnavailable(ctx *gin.Context) {
	New(ctx, errorx.NewCodeError(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable)))
}

func New(ctx *gin.Context, data interface{}) {
	resp := newResponse(ctx, data)
	resp.TraceId = fmt.Sprintf("%s", ctx.Value("traceId"))
	status := getHttpStatus(data)
	ctx.JSON(status, resp)
}

func getHttpStatus(data interface{}) int {
	switch value := data.(type) {

	case *errorx.CodeError:
		if value.Code >= 100 && value.Code <= 511 { // 使用标准的http状态码
			return value.Code
		}
	case error:
		return http.StatusInternalServerError
	}
	return http.StatusOK
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
