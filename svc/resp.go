package svc

import (
	"context"
	"fmt"
)

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceid"`
}

type Response struct {
	ctx context.Context
}

func NewResponse() *Response {
	return &Response{
		ctx: context.Background(),
	}
}

func (that *Response) WithContext(ctx context.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

func (that *Response) Success(data interface{}) Message {

	return Message{
		Message: "成功",
		Data:    data,
		TraceId: fmt.Sprintf("%s", that.ctx.Value("traceId")),
	}
}

func (that *Response) SuccessWithMessage(message string, data interface{}) Message {
	return Message{
		Message: message,
		Data:    data,
		TraceId: fmt.Sprintf("%s", that.ctx.Value("traceId")),
	}
}

func (that *Response) Error(message string) Message {
	return Message{
		Code:    1,
		Message: message,
		TraceId: fmt.Sprintf("%s", that.ctx.Value("traceId")),
	}
}

func (that *Response) ErrorWithCode(code int, message string) Message {
	return Message{
		Code:    code,
		Message: message,
		TraceId: fmt.Sprintf("%s", that.ctx.Value("traceId")),
	}
}
