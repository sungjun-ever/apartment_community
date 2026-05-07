package errUtils

type AppError struct {
	Err     error
	Status  int
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, status int, code string) *AppError {
	return &AppError{
		Err:     err,
		Status:  status,
		Code:    code,
		Message: GetErrorMessage(code),
	}
}
