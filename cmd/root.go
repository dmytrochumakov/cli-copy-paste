package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dmytrochumakov/cli-copy-paste/config"
	"github.com/dmytrochumakov/cli-copy-paste/db"
)

var usagePrefix = "Usage: cli-copy-paste "

func usageMessage(cmd Command) string {
	return usagePrefix + cmd.Usage()
}

func Execute() {
	if len(os.Args) < 2 {
		var messages []string
		for _, cmd := range GetAllCommands() {
			messages = append(messages, cmd.Usage())
		}

		fmt.Println(usagePrefix + strings.Join(messages, " "))
		os.Exit(1)
	}
	input := os.Args[1]

	db, err := db.OpenDB()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	cfg := config.Create(db)

	var cmdErr error

	if cmd, ok := ParseCommand(input); ok {
		switch cmd {
		case Paste:
			if len(os.Args) < 3 {
				fmt.Println(usageMessage(cmd))
			}
			value := os.Args[2]
			cmdErr = cfg.PasteTask(value)
		case CopyTaskFull:
			if len(os.Args) < 2 {
				fmt.Println(usageMessage(cmd))
			}
			cmdErr = CopyTaskFullFunc(cfg)
		case CopyTaskNumber:
			if len(os.Args) < 2 {
				fmt.Println(usageMessage(cmd))
			}
			cmdErr = CopyTaskNumberFunc(cfg)
		case FollowTaskLink:
			if len(os.Args) < 2 {
				fmt.Println(usageMessage(cmd))
			}
			cmdErr = FollowTaskLinkCmd(cfg)
		case Delete:
			if len(os.Args) < 3 {
				fmt.Println(usageMessage(cmd))
			}
			key := os.Args[2]
			cmdErr = cfg.Delete(key)
		case List:
			if len(os.Args) < 2 {
				fmt.Println(usageMessage(cmd))
			}
			cmdErr = cfg.List()
		}
	}
	if cmdErr != nil {
		fmt.Fprintln(os.Stderr, cmdErr)
		os.Exit(1)
	}

}
