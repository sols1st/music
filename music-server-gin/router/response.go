package router

type ResponseBody struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *ResponseBody {
	return &ResponseBody{
		Code:    200,
		Success: true,
		Message: "操作成功",
		Type:    "success",
		Data:    data,
	}
}

func SuccessWithMessage(message string, data interface{}) *ResponseBody {
	return &ResponseBody{
		Code:    200,
		Success: true,
		Message: message,
		Type:    "success",
		Data:    data,
	}
}

func Error(message string) *ResponseBody {
	return &ResponseBody{
		Code:    500,
		Success: false,
		Message: message,
		Type:    "error",
	}
}

func BadRequest(message string) *ResponseBody {
	return &ResponseBody{
		Code:    400,
		Success: false,
		Message: message,
		Type:    "error",
	}
}

// Unauthorized 未授权响应
func Unauthorized(message string) *ResponseBody {
	return &ResponseBody{
		Code:    401,
		Success: false,
		Message: message,
		Type:    "error",
	}
}

// Forbidden 禁止访问响应
func Forbidden(message string) *ResponseBody {
	return &ResponseBody{
		Code:    403,
		Success: false,
		Message: message,
		Type:    "error",
	}
}
