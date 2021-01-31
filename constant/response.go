package constant

type ResponseCode int

const (
	SUCCESS  ResponseCode = 0   //成功
	CODE_404 ResponseCode = 404 // 页面未找到
	CODE_500 ResponseCode = 500 // 服务器错误
)

var codeTextMap = map[ResponseCode]string{
	SUCCESS:  "成功",
	CODE_404: "页面不存在",
	CODE_500: "服务器内部错误",
}

func GetCodeText(code ResponseCode) string {
	if value, ok := codeTextMap[code]; ok {
		return value
	}
	return "Unkown code text"
}
