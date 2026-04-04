package dto

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func SuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}

func ErrorResponse(message string) Response[any] {
	return Response[any]{
		Success: false,
		Message: message,
	}
}
