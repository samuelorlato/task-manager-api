package errors

type HTTPError struct {
	StatusCode  int
	Err         error
	Description string
}

func (h *HTTPError) Error() string {
	return h.Err.Error()
}
