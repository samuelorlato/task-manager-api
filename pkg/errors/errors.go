package errors

func NewValidationError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  400,
		Err:         err,
		Description: "Validation error",
	}
}

func NewRepositoryError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  500,
		Err:         err,
		Description: "Repository error",
	}
}

func NewGenericError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  500,
		Err:         err,
		Description: "Generic error",
	}
}
