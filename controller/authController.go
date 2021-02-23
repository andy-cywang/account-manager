package controller

import (
	"account-manager/auth"
	"account-manager/middleware"
	"net/http"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac AuthController) Login(w http.ResponseWriter, r *http.Request) {
	username, password, err := middleware.ValidateLoginCredentials(r)
	if err != nil {
		middleware.WriteErrResponse(w, err, http.StatusBadRequest)
	}

	code, signedToken, err := auth.Login(username, password)
	if err != nil {
		middleware.WriteErrResponse(w, err, code)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		HttpOnly: true,
	})
	middleware.WriteJSONResponse(w, "Logged in", http.StatusOK)
}