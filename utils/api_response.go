package utils

type ApiResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewApiResponse(status string, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status: status,
		Message: message,
		Data: data,
	}
}