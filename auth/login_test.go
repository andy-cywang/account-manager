package auth

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	existUser := &Credentials{
		Username: "admin",
		Password: "admin",
	}

	wrongPassUser := &Credentials{
		Username: "admin",
		Password: "wrong password",
	}

	nonExistUser := &Credentials{
		Username: "I am not here",
		Password: "test",
	}

	code, _, err := Login(existUser.Username, existUser.Password)
	assert.NoError(t, err)
	assert.Equal(t, code, http.StatusOK)

	code, _, err = Login(wrongPassUser.Username, wrongPassUser.Password)
	assert.Equal(t, err.Error(), "auth: login: wrong password")
	assert.Equal(t, code, http.StatusUnauthorized)

	code, _, err = Login(nonExistUser.Username, nonExistUser.Password)
	assert.Equal(t, err.Error(), "auth: login: user doesn't exist")
	assert.Equal(t, code, http.StatusUnauthorized)
}