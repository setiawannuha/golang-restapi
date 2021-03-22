package middlewares

import (
	"net/http"

	"example.com/app/helpers"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		helpers.Log("Test log")
		next.ServeHTTP(rw, r)
	})
}
