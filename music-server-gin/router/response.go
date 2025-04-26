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
