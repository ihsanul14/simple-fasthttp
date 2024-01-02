package error

type Error struct {
	Code    int
	Message string
}

func NewError(code int, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}
