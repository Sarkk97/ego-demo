package middlewares

import (
	"ego/user/auth"
	"ego/user/response"
	"net/http"
)

//AuthenticationMiddleware sets the jwt authentication
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			response.Error(w, err.Error(), 401, map[string]string{"Content-Type": "application/json"})
			return
		}
		next(w, r)
	}
}
