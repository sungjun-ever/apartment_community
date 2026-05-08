package response

import "github.com/gin-gonic/gin"

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
	Total int64 `json:"total,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

func OK(c *gin.Context, status int, data interface{}, meta *Meta) {
	response := Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	}
	
	c.JSON(status, response)
}

func Fail(c *gin.Context, status int, code, message string) {
	response := Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}

	c.JSON(status, response)
}
