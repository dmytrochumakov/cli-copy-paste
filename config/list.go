package config

import (
	"fmt"

	"github.com/dmytrochumakov/cli-copy-paste/db"
)

func (cfg *Config) List() error {
	pairs, err := db.ListKeys(cfg.DB)
	if err != nil {
		return fmt.Errorf("Failed to list keys: %v", err)
	}
	if len(pairs) == 0 {
		fmt.Println("No values stored.")
		return nil
	}
	for k, v := range pairs {
		fmt.Printf("%s: %s\n", k, v)
	}
	return nil
}
