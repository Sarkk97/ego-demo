package response

import (
	"encoding/json"
	"net/http"
)

//Error returns error response
func Error(w http.ResponseWriter, errmsg interface{}, code int, headers map[string]string) {
	//Add headers
	for k, v := range headers {
		w.Header().Add(k, v)
	}
	// Set reposne code
	w.WriteHeader(code)

	err := map[string]interface{}{
		"status":  "failed",
		"message": errmsg,
	}

	json.NewEncoder(w).Encode(err)
}

//Success returns a success response
func Success(w http.ResponseWriter, data interface{}, code int, headers map[string]string) {
	for k, v := range headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(code)

	success := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	json.NewEncoder(w).Encode(success)
}
