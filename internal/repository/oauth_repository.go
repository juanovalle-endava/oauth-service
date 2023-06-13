package repository

import (
	"database/sql"
	"fmt"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"go.uber.org/fx"
	"log"
)

const (
	getAllTokens = "SELECT token FROM tokens"
	createToken  = "INSERT INTO tokens (token) VALUES ($1)"
)

type OAuthRepository interface {
	ListTokens() ([]string, error)
	SaveToken(string) error
}

type oAuthRepository struct {
	db config.Database
}

func (o *oAuthRepository) ListTokens() ([]string, error) {
	var tokens []string
	err := o.db.Select(&tokens, getAllTokens)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no registry of tokens")
			log.Fatalln(err)
			return tokens, err
		}
		log.Fatalln(err)
		return tokens, err
	}
	return tokens, err
}

func (o *oAuthRepository) SaveToken(token string) error {
	_, err := o.db.Exec(createToken, token)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func NewOAuthRepository(databaseConfig config.Database) OAuthRepository {
	return &oAuthRepository{databaseConfig}
}

var Module = fx.Provide(NewOAuthRepository)
