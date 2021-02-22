package util

type CustomError struct{
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}

func NewCustomError(msg string) CustomError {
	return CustomError{
		Message: msg,
	}
}