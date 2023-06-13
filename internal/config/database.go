package config

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

type Database struct {
	*sqlx.DB
}

func NewDatabase(config Config) Database {
	var err error

	connInfo := fmt.Sprintf("host=%s:%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.DbName)

	client, err := sqlx.Open("pgx", connInfo)
	if err != nil {
		fmt.Printf("Failed to open connection to database: %s", err.Error())
		panic(err)
	}
	if err = client.Ping(); err != nil {
		fmt.Printf("Failed to ping database: %s", err.Error())
		panic(err)
	}
	return Database{
		client,
	}
}

var DatabaseModule = fx.Provide(NewDatabase)
