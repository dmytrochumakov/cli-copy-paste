package cmd

import (
	"fmt"

	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/db"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
	"github.com/pkg/browser"
)

func FollowTaskLinkCmd(cfg *config.Config) error {
	val, err := db.GetValue(cfg.DB, feature.TASK_ID_KEY)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://dev.azure.com/competommc/Marketplace/_workitems/edit/%s", val)
	err = browser.OpenURL(url)
	if err != nil {
		return err
	}
	return nil
}
