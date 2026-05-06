package dto

import "github.com/gin-gonic/gin"

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

func OK(c *gin.Context, status *int, data interface{}, meta *Meta) {
	if status == nil {
		status = new(int)
		*status = 200
	}

	c.JSON(*status, Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func AbortWithError(c *gin.Context, status int, code, message string) {
	response := Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}
	c.AbortWithStatusJSON(status, response)
}

func Fail(c *gin.Context, status *int, code, message string) {
	if status == nil {
		status = new(int)
		*status = 400
	}

	response := Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}

	c.JSON(*status, response)
}
