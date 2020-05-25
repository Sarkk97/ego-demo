package respond

import (
	"encoding/json"
	"net/http"
)

//WithError writes an error message a ResponseWriter
func WithError(
	w http.ResponseWriter,
	errMsg interface{},
	code int,
	headers map[string]string) {

	for k, v := range headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(code)

	err := map[string]interface{}{
		"status": false,
		"error":  errMsg,
	}

	json.NewEncoder(w).Encode(err)
}

//WithSuccess writes a success message to a ResponseWriter
func WithSuccess(
	w http.ResponseWriter,
	data interface{},
	code int,
	headers map[string]string) {

	for k, v := range headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(code)

	success := map[string]interface{}{
		"status": true,
		"data":   data,
	}

	json.NewEncoder(w).Encode(success)
}
