package middlewares

import (
	"context"
	"ego/user/auth"
	"ego/user/response"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type ContextKey string

//AuthenticationMiddleware sets the jwt authentication
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			response.Error(w, err.Error(), 401, map[string]string{"Content-Type": "application/json"})
			return
		}
		//extract userid from token and add to request context
		tokenString, err := auth.ExtractToken(r)
		if err != nil {
			response.Error(w, err.Error(), 401, map[string]string{"Content-Type": "application/json"})
			return
		}
		userID, err := auth.GetIDFromAccessToken(tokenString)
		if err != nil {
			response.Error(w, err.Error(), 401, map[string]string{"Content-Type": "application/json"})
			return
		}
		ctx := context.WithValue(r.Context(), ContextKey("UserID"), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//TerminalLoggingHandler is a middleware that logs in apache format
func TerminalLoggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
