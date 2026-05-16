package errUtils

type AppError struct {
	Err    error
	Code   ErrorCode
	Status int
	Level  ErrorLevel
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, status int, code ErrorCode, level ErrorLevel) *AppError {
	appErr := &AppError{
		Err:    err,
		Code:   code,
		Status: status,
		Level:  level,
	}

	return appErr
}
