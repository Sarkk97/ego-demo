package middlewares

import (
	"ego/user/auth"
	"ego/user/response"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//AuthenticationMiddleware sets the jwt authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			response.Error(w, err.Error(), 401, map[string]string{"Content-Type": "application/json"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

//TerminalLoggingHandler is a middleware that logs in apache format
func TerminalLoggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
