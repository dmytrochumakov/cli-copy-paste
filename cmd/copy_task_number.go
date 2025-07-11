package cmd

import (
	"github.com/dmytrochumakov/cli-copy-paste/clipboard"
	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
)

func CopyTaskNumberFunc(cfg *config.Config) error {
	err := clipboard.Write(cfg, feature.TASK_ID_KEY)
	if err != nil {
		return err
	}
	return nil
}
