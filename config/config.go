package config

import (
	"database/sql"
)

type Config struct {
	DB *sql.DB
}

func Create(db *sql.DB) *Config {
	cfg := &Config{DB: db}

	return cfg
}
