package driver

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DB holds the database connection pool
type DB struct {
	Connection *pgxpool.Pool
}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := OpenDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.Config().MaxConns = maxOpenDbConn
	d.Config().MaxConnIdleTime = maxIdleDbConn
	d.Config().MaxConnLifetime = maxDbLifetime

	dbConn := &DB{Connection: d}

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// OpenDatabase opens the database for the application
func OpenDatabase(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
