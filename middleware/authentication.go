package middleware

import (
	"account-manager/auth"
	"net/http"
)

func Authenticate(handlerFunc func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("token")
		if clientToken == "" {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			return
		}

		jwtWrapper := auth.JWTWrapper{
			Issuer:    "AuthService",
		}

		_, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
	}
}