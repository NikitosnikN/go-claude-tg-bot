package errors

type ApplicationError struct {
	message string
}

func (a *ApplicationError) Error() string {
	return a.message
}

func NewApplicationError(message string) *ApplicationError {
	return &ApplicationError{
		message: message,
	}
}

var (
	DialogExpired = NewApplicationError("dialog expired")
)
