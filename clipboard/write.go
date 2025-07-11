package clipboard

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/dmytrochumakov/cli-copy-paste/config"
)

func Write(cfg *config.Config, value string) error {
	if err := clipboard.WriteAll(value); err != nil {
		return fmt.Errorf("Failed to copy to clipboard: %v", err)
	}
	fmt.Printf("Copied value for '%s' to clipboard\n", value)
	return nil
}
