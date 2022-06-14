package helper

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WrapperResponse(code int, success bool, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
}
