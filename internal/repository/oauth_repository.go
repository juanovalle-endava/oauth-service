package repository

import (
	"database/sql"
	"fmt"
	"github.com/juanovalle-endava/oauth-service/internal/config"
	"github.com/juanovalle-endava/oauth-service/internal/models"
	"go.uber.org/fx"
	"log"
)

const (
	getUserById  = "SELECT * FROM users WHERE user_id = $1"
	getAllTokens = "SELECT token FROM tokens"
	createToken  = "INSERT INTO tokens (token) VALUES ($1)"
)

type OAuthRepository interface {
	GetUserByUserId(username string) (models.User, error)
	ListTokens() ([]string, error)
	SaveToken(string) error
}

type oAuthRepository struct {
	db config.Database
}

func (o *oAuthRepository) GetUserByUserId(username string) (models.User, error) {
	var user models.User
	err := o.db.Get(&user, getUserById, username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no registry of tokens")
			log.Fatalln(err)
			return user, err
		}
		log.Fatalln(err)
		return user, err
	}
	return user, err
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
