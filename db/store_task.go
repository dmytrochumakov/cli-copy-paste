package db

import (
	"database/sql"
	"fmt"
	"github.com/dmytrochumakov/cli-copy-paste/feature"
)

func StoreTask(db *sql.DB, value string) error {
	if err := SetValue(db, feature.TASK_ID_KEY, value); err != nil {
		return fmt.Errorf("failed to save: %v", err)
	}
	fmt.Printf("Saved '%s'\n", value)
	return nil
}
