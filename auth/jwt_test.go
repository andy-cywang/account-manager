package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenGenerateAndValidation(t *testing.T) {
	jwtWrapper := JWTWrapper{
		Issuer: "AuthService",
	}

	generatedToken, err := jwtWrapper.GenerateToken("John")
	assert.NoError(t, err)

	claims, err := jwtWrapper.ValidateToken(generatedToken)
	assert.NoError(t, err)

	assert.Equal(t, "John", claims.Username)
}
