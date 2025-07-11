package cmd

import (
	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/db"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
)

func DeleteTaskCmd(cfg *config.Config) error {
	err := db.DeleteKey(cfg.DB, feature.TASK_ID_KEY)
	if err != nil {
		return err
	}
	return nil
}
