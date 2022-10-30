package svc

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

func (that *Response) Success(data interface{}) Message {

	return Message{
		Message: "成功",
		Data:    data,
	}
}

func (that *Response) SuccessWithMessage(message string, data interface{}) Message {
	return Message{
		Message: message,
		Data:    data,
	}
}

func (that *Response) Error(message string) Message {
	return Message{
		Code:    1,
		Message: message,
	}
}

func (that *Response) ErrorWithCode(code int, message string) Message {
	return Message{
		Code:    code,
		Message: message,
	}
}
