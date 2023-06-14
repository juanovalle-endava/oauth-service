package service

import (
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTokenCreator struct{}

func (m *MockTokenCreator) CreateToken(customerID string) (string, error) {
	return "", nil
}

func (m *MockTokenCreator) VerifyToken(token string) error {
	return nil
}

func TestListSigningKeys(t *testing.T) {
	tokenCreator := &MockTokenCreator{}
	oauthService := NewOAuthService(tokenCreator, config.Config{
		PublicKeyPath: "../mock/mock_jwtRS256.key.pub",
	})

	expectedResponse := models.KeyResponse{
		Keys: []models.Key{
			{
				Kty: "RSA",
				Use: "sig",
				Alg: "RS256",
				N:   "vGnUHnpUJPzTgXiIzsbcFHOA5zVkuRyhEed_qEEzgliJuHZ2FnOCbB3yW19u9rxMIDejbXhtIoQWDvcC1REkJqQzA8CEXWMx1zASCGsHOkq2ogv0mBL-hE2ZiORnwcEDSz9LHgqpXiIFflnBFxFkPmK-hUnnI11DEg5wGOmZMHs=",
			},
		},
	}

	response, err := oauthService.ListSigningKeys()

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestListSigningKeysNoKeysFound(t *testing.T) {
	tokenCreator := &MockTokenCreator{}
	oauthService := NewOAuthService(tokenCreator, config.Config{
		PublicKeyPath: "../mock/no_key.key.pub",
	})

	response, err := oauthService.ListSigningKeys()

	assert.Error(t, err)
	assert.Empty(t, response)
}
