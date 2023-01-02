package response

import (
	"context"
	"fmt"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceid"`
}

func New(ctx context.Context, code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		TraceId: fmt.Sprintf("%s", ctx.Value("traceId")),
	}
}

func NewDefault(ctx context.Context, data interface{}) *Response {
	return New(ctx, 0, "成功", data)
}

func NewDefaultError(ctx context.Context, err error) *Response {
	return NewCodeError(ctx, 1, err)
}
func NewCodeError(ctx context.Context, code int, err error) *Response {
	return New(ctx, code, err.Error(), nil)
}
