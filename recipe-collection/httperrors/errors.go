package httperrors

//HTTPError is an error that occured during request processing
type HTTPError struct {
	Status  int
	Message string
}

//ServerError holds the message displayed to the client
// on on 5xx http error
const ServerError = "A server error occurred"

//NewHTTPError returns an instance of HTTPError
func NewHTTPError(message string, status int) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
	}
}

//Error returns HTTPError string
func (error HTTPError) Error() string {
	return error.Message
}
