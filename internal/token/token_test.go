package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTokenCreator(t *testing.T) {
	tokenCreator := NewTokenCreator(config.Config{
		PrivateKeyPath:       "../mock/mock_jwtRS256.key",
		PublicKeyPath:        "../mock/mock_jwtRS256.key.pub",
		TokenDefaultDuration: "24h",
	})

	tkn, err := tokenCreator.CreateToken("endava")
	assert.NoError(t, err)
	assert.NotEmpty(t, tkn)

	tokenPayload, err := tokenCreator.VerifyToken(tkn)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenPayload)
}

func TestExpiredJWTToken(t *testing.T) {
	tokenCreator := NewTokenCreator(config.Config{
		PrivateKeyPath:       "../mock/mock_jwtRS256.key",
		PublicKeyPath:        "../mock/mock_jwtRS256.key.pub",
		TokenDefaultDuration: "-24h",
	})

	tkn, err := tokenCreator.CreateToken("endava")
	assert.NoError(t, err)
	assert.NotEmpty(t, tkn)

	tokenPayload, err := tokenCreator.VerifyToken(tkn)
	assert.Error(t, err)
	assert.Empty(t, tokenPayload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {

	tokenCreator := NewTokenCreator(config.Config{
		PrivateKeyPath:       "../mock/mock_jwtRS256.key",
		PublicKeyPath:        "../mock/mock_jwtRS256.key.pub",
		TokenDefaultDuration: "-24h",
	})

	payload := models.TokenPayload{
		ClientId:  "n/a",
		IssuedAt:  time.Time{},
		ExpiredAt: time.Time{},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tkn, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	tokenPayload, err := tokenCreator.VerifyToken(tkn)
	assert.Error(t, err)
	assert.Empty(t, tokenPayload)
}

func TestInvalidJWTSignedSecret(t *testing.T) {
	tokenCreator := NewTokenCreator(config.Config{
		PrivateKeyPath:       "../mock/mock_jwtRS256.key",
		PublicKeyPath:        "../mock/mock_jwtRS256.key.pub",
		TokenDefaultDuration: "-24h",
	})

	payload := models.TokenPayload{
		ClientId:  "n/a",
		IssuedAt:  time.Time{},
		ExpiredAt: time.Time{},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tkn, err := jwtToken.SignedString([]byte("testPrivKey"))
	assert.NoError(t, err)

	tokenPayload, err := tokenCreator.VerifyToken(tkn)
	assert.Error(t, err)
	assert.Empty(t, tokenPayload)
}
