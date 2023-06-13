package token

import (
	"fmt"
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
	VerifyToken(token string) error
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

func (t *token) VerifyToken(token string) error {
	pubKeyBytes, err := os.ReadFile(t.config.PublicKeyPath)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return pubKey, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &models.TokenPayload{}, keyFunc)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if !jwtToken.Valid {
		err := fmt.Errorf("token signature is invalid")
		log.Fatalln(err)
		return err
	}

	return nil
}

func NewTokenCreator(config config.Config) Token {
	return &token{config: config}
}

var Module = fx.Provide(NewTokenCreator)
