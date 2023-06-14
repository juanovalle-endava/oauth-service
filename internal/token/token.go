package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"github.com/juanovalle-endava/oauth-service/internal/utils"
	"go.uber.org/fx"
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
	privKey, err := utils.GetPrivateKey(t.config.PrivateKeyPath)
	if err != nil {
		return "", err
	}

	duration, err := time.ParseDuration(t.config.TokenDefaultDuration)
	if err != nil {
		return "", err
	}

	payload := &models.TokenPayload{
		ClientId:  clientId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	tkn, err := jwtToken.SignedString(privKey)
	if err != nil {
		return "", err
	}

	return tkn, err
}

func (t *token) VerifyToken(token string) error {
	pubKey, err := utils.GetPublicKey(t.config.PublicKeyPath)
	if err != nil {
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
		return err
	}

	if !jwtToken.Valid {
		err = fmt.Errorf("token signature is invalid")
		return err
	}

	return nil
}

func NewTokenCreator(config config.Config) Token {
	return &token{config: config}
}

var Module = fx.Provide(NewTokenCreator)
