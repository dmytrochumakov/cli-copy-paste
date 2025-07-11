package cmd

import (
	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
	"github.com/dmytrochumakov/cli-copy-paste/db"
	"fmt"
)

func PasteTaskCmd(cfg *config.Config, value string) error {
	if err := db.SetValue(cfg.DB, feature.TASK_ID_KEY, value); err != nil {
		return fmt.Errorf("failed to save: %v", err)
	}
	fmt.Printf("Saved '%s'\n", value)

	return nil
}
