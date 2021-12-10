package helpers

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func RespApi(code int, status string, data interface{}) ApiResponse {
	return ApiResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}