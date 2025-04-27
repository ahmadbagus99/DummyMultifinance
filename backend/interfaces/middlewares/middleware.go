package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	handlers "DummyMultifinance/interfaces/handlers"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("SECRET_KEY"))

func TokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handlers.GeneralResponse(w, http.StatusBadRequest, handlers.Unauthorized, "Authorization header is missing", nil)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			handlers.GeneralResponse(w, http.StatusBadRequest, handlers.Unauthorized, "Token is missing", nil)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return signingKey, nil
		})

		if err != nil || !token.Valid {
			handlers.GeneralResponse(w, http.StatusBadRequest, handlers.Unauthorized, "Invalid token", nil)
			return
		}

		next(w, r)
	}
}
