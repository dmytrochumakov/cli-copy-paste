package config

import (
	"fmt"

	"github.com/dmytrochumakov/cli-copy-paste/db"
)

func (cfg *Config) Delete(key string) error {
	_, err := db.GetValue(cfg.DB, key)
	if err != nil {
		return fmt.Errorf("Key '%s' not found", key)
	}
	if err := db.DeleteKey(cfg.DB, key); err != nil {
		return fmt.Errorf("Failed to delete: %v", err)
	}
	fmt.Printf("Deleted '%s'\n", key)
	return nil
}
