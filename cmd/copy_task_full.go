package cmd

import (

	"github.com/dmytrochumakov/cli-copy-paste/clipboard"
	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/db"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
)

func CopyTaskFullFunc(cfg *config.Config) error {
	val, err := db.GetValue(cfg.DB,feature.TASK_ID_KEY)
	if err != nil {
		return err
	}
	val = "#" + val + " " + "-" + " "
	err = clipboard.Write(cfg, val)
	if err != nil {
		return err
	}
	return nil
}
