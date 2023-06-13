package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

type Token interface {
	CreateToken(customerID string) (string, error)
	VerifyToken(token string) (*models.TokenPayload, error)
}

type token struct {
	config config.Config
}

func (t *token) CreateToken(clientId string) (string, error) {
	privKeyBytes, err := os.ReadFile(t.config.PrivateKeyPath)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	duration, err := time.ParseDuration(t.config.TokenDefaultDuration)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	payload := &models.TokenPayload{
		ClientId:  clientId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	token, err := jwtToken.SignedString(privKey)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return token, err
}

func (t *token) VerifyToken(token string) (*models.TokenPayload, error) {
	//TODO implement me
	panic("implement me")
}

func NewTokenCreator(config config.Config) Token {
	return &token{config: config}
}

var Module = fx.Provide(NewTokenCreator)
