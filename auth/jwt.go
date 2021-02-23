package auth

import (
	"account-manager/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "secret_key"
	ExpirationHours = 1
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTWrapper struct {
	Issuer    string
}

func (j *JWTWrapper) GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(ExpirationHours)).Unix(),
			Issuer: j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", util.NewCustomError("auth: generate token: sign token: " + err.Error())
	}

	return signedToken, nil
}

func (j *JWTWrapper) ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		return nil, util.NewCustomError("auth: validate token: parse claims: " + err.Error())
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, util.NewCustomError("auth: validate token: get claims: " + err.Error())
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, util.NewCustomError("auth: validate token: JWT is expired: " + err.Error())
	}

	return claims, nil
}