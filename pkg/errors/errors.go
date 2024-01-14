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

func NewJWTInvalidError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  401,
		Err:         err,
		Description: "Invalid JWT",
	}
}

func NewJWTGenerationError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  500,
		Err:         err,
		Description: "Error while generating JWT",
	}
}

func NewAuthorizationError(err error) *HTTPError {
	return &HTTPError{
		StatusCode:  401,
		Err:         err,
		Description: "You must be authenticated",
	}
}
