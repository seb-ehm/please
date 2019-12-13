package rules

import (
	"please/pkg/command"
)

//AllRules returns all commands known to please
func AllRules() []command.Details {

	allRules := []command.Details{mv{}, ls{}}

	return allRules
}

type mv struct{}

func (mv) Name() string {
	return "mv"
}

func (mv) Misspellings() []string {
	return []string{"move", "mve"}
}

func (mv) Correct(command command.Command) command.Command {
	return command
}

type ls struct{}

func (ls) Name() string {
	return "ls"
}

func (ls) Misspellings() []string {
	return []string{}
}

func (ls) Correct(command command.Command) command.Command {
	return command
}
