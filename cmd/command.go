package cmd

import "strings"

type Command int

const (
	PasteTask = iota
	CopyTaskNumber
	CopyTaskFull
	FollowTaskLink
	Delete
	List
)

type CommandPair struct {
	Long  string
	Short string
}

var commandMap = map[Command]CommandPair{
	PasteTask:      {"paste-task", "pt"},
	CopyTaskNumber: {"copy-task-number", "ctn"},
	CopyTaskFull:   {"copy-task-full", "ctf"},
	FollowTaskLink: {"follow-task-link", "f"},
	Delete:         {"delete", "d"},
	List:           {"list", "l"},
}

func ParseCommand(s string) (Command, bool) {
	for cmd, pair := range commandMap {
		if s == pair.Long || s == pair.Short {
			return cmd, true
		}
	}
	return 0, false
}

func (cmd Command) Usage() string {
	cmdPair := commandMap[cmd]
	cmdStrs := make([]string, 2)
	cmdStrs = append(cmdStrs, cmdPair.Long)
	cmdStrs = append(cmdStrs, cmdPair.Short)
	res := strings.Join(cmdStrs, "|")

	switch cmd {
	case PasteTask:
		return res + " <task_number>"
	case Delete:
		return res + " <key>"
	}
	return res
}

func GetAllCommands() []Command {
	return []Command{List, PasteTask, CopyTaskNumber, CopyTaskFull, Delete}
}
