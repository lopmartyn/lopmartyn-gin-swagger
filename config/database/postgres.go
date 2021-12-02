package database

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// PostgresConfig persists the config for our PostgreSQL databases connection
type PostgresConfig struct {
	URL      string `env:"DATABASE_URL"` // DATABASE_URL will be used in preference if it exists
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"65432"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
}

var PGConnection *bun.DB

// CreateDatabaseConnection create database connection via BunDB
func CreateDatabaseConnection(pg PostgresConfig) {
	// Open a PostgreSQL databases.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pg.User, pg.Password, pg.Host, pg.Port, pg.Database)
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// Create a poc-boilerplate db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())
	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	PGConnection = db
}
