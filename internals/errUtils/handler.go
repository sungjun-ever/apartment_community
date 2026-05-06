package errUtils

type AppError struct {
	Err     error
	Message string
	Status  int
	Code    string
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
