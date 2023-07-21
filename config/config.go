package config

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dsn = ""

func SetDSN(d string) {
	dsn = d
}

type Config struct {
	Postgres Postgres
	Context  context.Context
}

type Postgres struct {
	DSN                  string
	Conn                 *sqlx.DB
	RefreshAwaitDuration time.Duration
	RefreshCheckPause    time.Duration
}

func MustConfig(ctx context.Context) Config {
	cfg, err := createConfig(ctx)
	if err != nil {
		panic(err)
	}

	return cfg
}

func createConfig(ctx context.Context) (Config, error) {
	cfg := getConfig(ctx)

	var err error
	cfg.Postgres.Conn, err = sqlx.Open("postgres", cfg.Postgres.DSN)
	cfg.Postgres.DSN = ""

	if err == nil {
		err = cfg.Postgres.Conn.Ping()
	}

	return cfg, err
}

func getConfig(ctx context.Context) Config {
	cfg := Config{
		Context: ctx,
		Postgres: Postgres{
			DSN:                  dsn,
			RefreshCheckPause:    2000000000,
			RefreshAwaitDuration: 2000000000,
		},
	}

	return cfg
}
