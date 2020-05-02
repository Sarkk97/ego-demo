package httperror

//ServerError is a custom message for HTTP 5XX errors
const ServerError = "An error occured but it was not your fault"

//HTTPError is an HTTP error
type HTTPError struct {
	Message string
	Code    int
}

//Error is a method implemented by all error (go builtin) types
func (httpError HTTPError) Error() string {
	return httpError.Message
}
