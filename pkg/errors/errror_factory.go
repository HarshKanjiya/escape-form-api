package errors

import "net/http"

func BadRequest(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Message:    msg,
	}
}

func Unauthorized(msg string) *AppError {
	if msg == "" {
		msg = "Unauthorized access"
	} else {
		msg = "You don't have access to this " + msg
	}
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
		Message:    msg,
	}
}

func PaymentRequired(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusPaymentRequired, // 402
		Code:       "PAYMENT_REQUIRED",
		Message:    msg,
	}
}

func NotFound(name string) *AppError {
	if name == "" {
		name = "Resource"
	}
	return &AppError{
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Message:    name + " not found",
	}
}

func Internal(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_ERROR",
		Message:    "Something went wrong",
		Err:        err,
	}
}
