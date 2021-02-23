package auth

import (
	"account-manager/util"
	"net/http"
)

var users = map[string]string{
	"admin": "admin",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(username, password string) (code int, signedToken string, err error) {
	if v, ok := users[username]; !ok {
		return http.StatusUnauthorized, "", util.NewCustomError("auth: login: user doesn't exist")
	} else {
		if password != v {
			return http.StatusUnauthorized, "", util.NewCustomError("auth: login: wrong password")
		}
	}

	jwtWrapper := JWTWrapper {
		Issuer:    "AuthService",
	}

	signedToken, err = jwtWrapper.GenerateToken(username)
	if err != nil {
		return http.StatusInternalServerError, "", util.NewCustomError("auth: generator token: " +  err.Error())
	}

	return http.StatusOK, signedToken, nil
}