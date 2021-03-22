package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"example.com/app/helpers"
	"github.com/dgrijalva/jwt-go"
)

func Authorization(next http.Handler) http.Handler {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SECRET = os.Getenv("JWT_SECRET")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		res := helpers.Response{ResponseWritter: rw}
		bearerToken := r.Header.Get("Authorization")
		tokenString := strings.Replace(bearerToken, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}
			return []byte(JWT_SECRET), nil
		})
		if err != nil {
			res.SendError(nil, "Error Unauthorized")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			res.SendError(nil, "Error Unauthorized")
			return
		}
		fmt.Println("Claims", claims)
		next.ServeHTTP(rw, r)
	})
}
