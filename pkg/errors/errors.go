package errors

// Application-level error
type AppError struct {
	StatusCode int
	Code       string
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	return e.Message
}
