package dbrepo

import (
	"github.com/guicattani/handy-folk-finder/internal/config"
	"github.com/guicattani/handy-folk-finder/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *pgxpool.Pool
}

// Creates Postgres struct as a DB repo
func NewPostgresRepo(conn *pgxpool.Pool, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
