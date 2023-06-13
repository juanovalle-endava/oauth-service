package service

import (
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"github.com/juanovalle-endava/oauth-service/internal/repository"
	"go.uber.org/fx"
)

type OAuthService interface {
	GetUserByUserId(username string) (models.User, error)
	ListTokens() ([]string, error)
	SaveToken(string) error
}

type oAuthService struct {
	repository.OAuthRepository
}

func NewOAuthService(authRepository repository.OAuthRepository) OAuthService {
	return &oAuthService{
		authRepository,
	}
}

var Module = fx.Provide(NewOAuthService)
