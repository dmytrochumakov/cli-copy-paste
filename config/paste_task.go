package config

import (
	"github.com/dmytrochumakov/cli-copy-paste/db"
)

func (cfg *Config) PasteTask(value string) error {
	db.StoreTask(cfg.DB, value)
	return nil
}
