package models

type ApiError struct {
	Status int
	Message string
}

func (a ApiError) Error() string {
	return a.Message
}

func FromError(err error, status int) *ApiError {
	return &ApiError{
		Status: status,
		Message: err.Error(),
	}
}
