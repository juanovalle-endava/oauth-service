package service

import (
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"github.com/juanovalle-endava/oauth-service/internal/token"
	"github.com/juanovalle-endava/oauth-service/internal/utils"
	"go.uber.org/fx"
)

type OAuthService interface {
	CreateToken(customerID string) (string, error)
	VerifyToken(token string) error
	ListSigningKeys() (models.KeyResponse, error)
}

type oAuthService struct {
	token.Token
	env config.Config
}

func (o *oAuthService) ListSigningKeys() (models.KeyResponse, error) {
	var keys []models.Key

	pubkey, err := utils.GetPublicKey(o.env.PublicKeyPath)
	if err != nil {
		return models.KeyResponse{}, err
	}

	keys = append(keys, models.Key{
		Kty: "RSA",
		Use: "sig",
		Alg: "RS256",
		N:   utils.EncodeToBase64URL(pubkey.N.Bytes()),
	})

	return models.KeyResponse{
		Keys: keys,
	}, nil
}

func NewOAuthService(tokenCreator token.Token, config config.Config) OAuthService {
	return &oAuthService{tokenCreator, config}
}

var Module = fx.Provide(NewOAuthService)
